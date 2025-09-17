package models

import (
	"gorm.io/gorm"
	"time"
)

type Reimbursement struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Amount      float64        `json:"amount"`
	CategoryID  uint           `json:"category_id"`
	Status      string         `json:"status"` // pending, approved, rejected
	SubmittedAt time.Time      `json:"submitted_at"`
	ApprovedAt  *time.Time     `json:"approved_at"`
	FilePath    string         `json:"file_path"`
	UserEmail   string         `json:"user_email"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
