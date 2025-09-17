package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"reimbursement-service-go/config"
	"reimbursement-service-go/models"
	"reimbursement-service-go/utils"
)

func CreateReimbursement(c *gin.Context) {
	email := c.GetString("user_email")
	role := c.GetString("user_role")
	if role != "EMPLOYEE" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only employees can submit"})
		return
	}

	var req struct {
		Title       string  `form:"title" binding:"required"`
		Description string  `form:"description"`
		Amount      float64 `form:"amount" binding:"required"`
		CategoryID  uint    `form:"category_id" binding:"required"`
	}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// validasi limit per bulan
	var category models.Category
	if err := config.DB.First(&category, req.CategoryID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category not found"})
		return
	}

	var total float64
	startMonth := time.Now().UTC().Truncate(24*time.Hour).AddDate(0, 0, -time.Now().Day()+1)
	endMonth := startMonth.AddDate(0, 1, 0)
	config.DB.Model(&models.Reimbursement{}).
		Where("user_email = ? AND category_id = ? AND submitted_at BETWEEN ? AND ? AND status <> ?", email, req.CategoryID, startMonth, endMonth, "rejected").
		Select("COALESCE(SUM(amount),0)").Row().Scan(&total)

	if total+req.Amount > category.LimitPerMonth {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Exceeds monthly category limit"})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}
	if file.Size > 2*1024*1024 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File max 2MB"})
		return
	}

	filePath, err := utils.SaveUploadedFile(file, "./uploads")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	reimbursement := models.Reimbursement{
		Title:       req.Title,
		Description: req.Description,
		Amount:      req.Amount,
		CategoryID:  req.CategoryID,
		Status:      "pending",
		SubmittedAt: time.Now(),
		UserEmail:   email,
		FilePath:    filePath,
	}

	if err := config.DB.Create(&reimbursement).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// log activity
	config.DB.Create(&models.Log{
		Action:    fmt.Sprintf("New reimbursement submitted: %s", reimbursement.Title),
		UserEmail: email,
		CreatedAt: time.Now(),
	})

	c.JSON(http.StatusCreated, reimbursement)
}

// ================= Manager/Admin =================
func ApproveReimbursement(c *gin.Context) {
	role := c.GetString("user_role")
	if role != "MANAGER" && role != "ADMIN" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only manager/admin can approve"})
		return
	}

	id := c.Param("id")
	var reimbursement models.Reimbursement
	if err := config.DB.First(&reimbursement, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reimbursement not found"})
		return
	}

	reimbursement.Status = "approved"
	now := time.Now()
	reimbursement.ApprovedAt = &now

	config.DB.Save(&reimbursement)

	// log activity
	config.DB.Create(&models.Log{
		Action:    fmt.Sprintf("Reimbursement approved: %s", reimbursement.Title),
		UserEmail: c.GetString("user_email"),
		CreatedAt: time.Now(),
	})

	c.JSON(http.StatusOK, reimbursement)
}

func RejectReimbursement(c *gin.Context) {
	role := c.GetString("user_role")
	if role != "MANAGER" && role != "ADMIN" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only manager/admin can reject"})
		return
	}

	id := c.Param("id")
	var reimbursement models.Reimbursement
	if err := config.DB.First(&reimbursement, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reimbursement not found"})
		return
	}

	reimbursement.Status = "rejected"
	now := time.Now()
	reimbursement.ApprovedAt = &now
	config.DB.Save(&reimbursement)

	// log activity
	config.DB.Create(&models.Log{
		Action:    fmt.Sprintf("Reimbursement rejected: %s", reimbursement.Title),
		UserEmail: c.GetString("user_email"),
		CreatedAt: time.Now(),
	})

	c.JSON(http.StatusOK, reimbursement)
}

// GET reimbursements, employee sees own, manager/admin sees all
func GetAllReimbursements(c *gin.Context) {
	role := c.GetString("user_role")
	email := c.GetString("user_email")

	var reimbursements []models.Reimbursement
	if role == "EMPLOYEE" {
		config.DB.Where("user_email = ?", email).Find(&reimbursements)
	} else if role == "MANAGER" || role == "ADMIN" {
		config.DB.Unscoped().Find(&reimbursements) // termasuk soft-deleted
	}

	c.JSON(http.StatusOK, reimbursements)
}

// Soft delete reimbursement
func DeleteReimbursement(c *gin.Context) {
	role := c.GetString("user_role")
	if role != "ADMIN" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only admin can delete"})
		return
	}

	id := c.Param("id")
	if err := config.DB.Delete(&models.Reimbursement{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// log
	config.DB.Create(&models.Log{
		Action:    fmt.Sprintf("Reimbursement soft-deleted ID: %s", id),
		UserEmail: c.GetString("user_email"),
		CreatedAt: time.Now(),
	})

	c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}
