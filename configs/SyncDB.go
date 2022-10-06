package configs

import "go-gorm-jwt/models"

func SyncDB() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Book{})
}