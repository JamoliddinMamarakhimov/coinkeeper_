package models

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	FullName  string    `json:"full_name"`
	Username  string    `json:"username" gorm:"unique"`
	Password  string    `json:"password" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type SwagUser struct {
	FullName string `json:"full_name"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password" gorm:"not null"`
}

type SignInInput struct {
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password" gorm:"not null"`
}
