package routes

import (
	"github.com/gin-gonic/gin"
	"reimbursement-service-go/controller"
	"reimbursement-service-go/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// JWT group
	reimbursement := r.Group("/api/v1/reimbursements")
	reimbursement.Use(middleware.JWTAuthMiddleware())
	{
		reimbursement.POST("/", controllers.CreateReimbursement)
		reimbursement.GET("/", controllers.GetAllReimbursements)
		reimbursement.PUT("/approve/:id", controllers.ApproveReimbursement)
		reimbursement.PUT("/reject/:id", controllers.RejectReimbursement)
		reimbursement.DELETE("/:id", controllers.DeleteReimbursement)
	}

	// Handle 404
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"error": "Resource not found"})
	})

	return r
}
