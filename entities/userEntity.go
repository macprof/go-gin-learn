package entities

import "gorm.io/gorm"

type UserEntity struct {
	gorm.Model

	ID        int    `gorm:"type:int;primaryKey" json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
	Platform  string `json:"platform"`
	// Role      string `json:"role"`
	// Timestamp
}

func (UserEntity) TableName() string {
	return "table_users"
}
