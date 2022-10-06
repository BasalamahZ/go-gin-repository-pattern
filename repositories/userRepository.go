package repositories

import (
	"go-gorm-jwt/configs"
	"go-gorm-jwt/models"
)

type UserRepository interface {
	FindAll() ([]models.User, error)
	FindByID(userID string) (models.User, error)
	Delete(userID string) (models.User, error)
}

type userRepository struct {
}

// FindAll implements UserRepository
func (*userRepository) FindAll() ([]models.User, error) {
	var user []models.User
	db := configs.DB.Find(&user)
	if db != nil {
		return user, db.Error
	}
	return user, nil
}

// FindByID implements UserRepository
func (*userRepository) FindByID(userID string) (models.User, error) {
	var user models.User
	db := configs.DB.First(&user, userID)
	if db != nil {
		return user, db.Error
	}
	return user, nil
}

// Delete implements UserRepository
func (*userRepository) Delete(userID string) (models.User, error) {
	var user models.User
	db := configs.DB.Delete(userID)
	if db != nil {
		return user, db.Error
	}
	return user, nil
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}
