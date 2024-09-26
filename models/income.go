package models

import "time"

type Income struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	Description string    `json:"description"`
	Amount      float32   `json:"amount"`
	User        User      `json:"-" gorm:"foreignKey:UserID;references:ID"`
	UserID      uint      `json:"-"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
	IsDeleted   bool      `json:"-" gorm:"default:false"`
}
