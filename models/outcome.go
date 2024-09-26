package models

import "time"

type Outcome struct {
	ID          uint            `json:"id" gorm:"primary_key"`
	Description string          `json:"description"`
	Category    OutcomeCategory `json:"-" gorm:"foreignKey:CategoryID;references:ID"`
	CategoryID  uint            `json:"category_id"`
	Amount      float32         `json:"amount" gorm:"not null"`
	User        User            `json:"-" gorm:"foreignKey:UserID;references:ID"`
	UserID      uint            `json:"-"`
	CreatedAt   time.Time       `json:"-"`
	UpdatedAt   time.Time       `json:"-"`
	IsDeleted   bool            `json:"-" gorm:"default:false"`
}

type OutcomeCategory struct {
	ID    uint64 `json:"id" gorm:"primaryKey"`
	Title string `json:"title" gorm:"not null"`
}
