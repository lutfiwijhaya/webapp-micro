package models

type Category struct {
	ID             uint            `gorm:"primaryKey" json:"id"`
	Name           string          `json:"name"`
	LimitPerMonth  float64         `json:"limit_per_month"`
	Reimbursements []Reimbursement `gorm:"foreignKey:CategoryID" json:"-"`
}
