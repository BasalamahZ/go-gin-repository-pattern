package services

import (
	"go-gorm-jwt/models"
	"go-gorm-jwt/repositories"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(registerRequest models.User) (models.User, error)
	Login(loginRequest models.LoginRequest) (models.User, error)
}

type authService struct {
	authRepository repositories.AuthRepository
}

// Register implements AuthService
func (as *authService) Register(registerRequest models.User) (models.User, error) {
	user := models.User{}
	user.Name = registerRequest.Name
	user.Email = registerRequest.Email
	hash, err := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), 10)
	if err != nil {
		return user, err
	}
	user.Password = string(hash)

	newUser, err := as.authRepository.Create(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

// Login implements AuthService
func (as *authService) Login(loginRequest models.LoginRequest) (models.User, error) {
	user, err := as.authRepository.FindByEmail(loginRequest.Email)
	if err != nil {
		return user, err
	}
	
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func NewAuthService(authRepository *repositories.AuthRepository) AuthService {
	return &authService{
		authRepository: *authRepository,
	}
}
