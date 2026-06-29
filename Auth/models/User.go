package models

import (
	"time"
)
type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Password  string    `json:"password" gorm:"not null"` // "-" hides from JSON
	Role      string    `json:"role" gorm:"default:user"` // admin, user
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserResponse struct {
    ID    uint   
    Name  string 
    Email string 
    Role  string 
	Token string 
}
