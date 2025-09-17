package models

import "time"

type Log struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Action    string    `json:"action"`
	UserEmail string    `json:"user_email"`
	CreatedAt time.Time `json:"created_at"`
}
