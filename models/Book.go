package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title       string
	Description string
	Author      string
	Price       int
	Rating      float32
}
