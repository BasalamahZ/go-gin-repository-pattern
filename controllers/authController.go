package controllers

import (
	"go-gorm-jwt/helpers"
	"go-gorm-jwt/middleware"
	"go-gorm-jwt/models"
	"go-gorm-jwt/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(authService *services.AuthService) AuthController {
	return AuthController{
		authService: *authService,
	}
}

func (ac *AuthController) Register(c *gin.Context) {
	// get the email/pass req body
	registerRequest := new(models.User)

	err := c.BindJSON(&registerRequest)
	if err != nil {
		errors := helpers.FormatError(err)
		errorMassage := gin.H{"errors": errors}
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errorMassage,
		})
		return
	}

	user, err := ac.authService.Register(*registerRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed To Register",
		})
		return
	}
	//response
	c.JSON(http.StatusOK, gin.H{
		"message": "Registration Is Successfully",
		"status":  true,
		"data":    user,
	})
}

func (ac *AuthController) Login(c *gin.Context) {
	// get email password off req body
	var loginRequest models.LoginRequest
	err := c.BindJSON(&loginRequest)
	if err != nil {
		errors := helpers.FormatError(err)
		errorMassage := gin.H{"errors": errors}
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errorMassage,
		})
		return
	}

	user, err := ac.authService.Login(loginRequest)
	// if user.ID == 0 {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"message": "Invalid Email or Password",
	// 	})
	// 	return
	// }
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed To Login",
		})
		return
	}
	userID := user.ID
	// generate jwt token
	token, err := middleware.GenerateToken(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed To Generate Token",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", token, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login Is Successfully",
		"token":   token,
		"status":  true,
	})

}
