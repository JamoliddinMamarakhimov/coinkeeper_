package models

import "time"

type Expense struct {
	ID          uint    `json:"id" gorm:"primary_key"`
	Amount      float32 `json:"amount"`
	Description string  `json:"description"`

	Card   Card `json:"-" gorm:"foreignKey:CardID;references:ID"`
	CardID uint `json:"card_id"`

	Category   OutcomeCategory `json:"-" gorm:"foreignKey:CategoryID;references:ID"`
	CategoryID uint            `json:"category_id"`

	User   User `json:"-" gorm:"foreignKey:UserID;references:ID"`
	UserID uint `json:"user_id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsDeleted bool      `json:"is_deleted"`
}
