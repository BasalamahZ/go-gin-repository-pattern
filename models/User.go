package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `binding:"required"`
	Email    string `gorm:"unique" binding:"required,email"`
	Password string `binding:"required,min=6"`
}

type LoginRequest struct {
	Email    string `binding:"required,email"`
	Password string `binding:"required,min=6"`
}
