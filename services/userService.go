package services

import (
	"go-gorm-jwt/models"
	"go-gorm-jwt/repositories"
)

type UserService interface {
	FindAll() ([]models.User, error)
	FindByID(userID string) (models.User, error)
	Delete(userID string) (models.User, error)
}

type userService struct {
	userRepository repositories.UserRepository
}

// FindAll implements UserService
func (us *userService) FindAll() ([]models.User, error) {
	user, err := us.userRepository.FindAll()
	if err != nil {
		return user, err
	}
	return user, nil
}

// FindByID implements UserService
func (us *userService) FindByID(userID string) (models.User, error) {
	user, err := us.userRepository.FindByID(userID)
	if err != nil {
		return user, err
	}
	return user, nil
}

// Delete implements UserService
func (us *userService) Delete(userID string) (models.User, error) {
	user, err := us.userRepository.Delete(userID)
	if err != nil {
		return user, err
	}
	return user, nil
}

func NewUserService(userRepository *repositories.UserRepository) UserService {
	return &userService{
		userRepository: *userRepository,
	}
}
