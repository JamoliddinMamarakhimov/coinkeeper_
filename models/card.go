package models

import "time"

type Card struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	CardNumber  string    `json:"card_number"`
	Balance     float32   `json:"balance" gorm:"not null"`
	Description string    `json:"description"`
	User        User      `json:"-" gorm:"foreignKey:UserID;references:ID"` // Внешний ключ к User
	UserID      uint      `json:"user_id"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
	IsDeleted   bool      `json:"-" gorm:"default:false"`
}
