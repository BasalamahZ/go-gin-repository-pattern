package controllers

import (
	"go-gorm-jwt/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService *services.UserService) UserController {
	return UserController{
		userService: *userService,
	}
}

func (uc *UserController) FindAll(c *gin.Context) {
	user, err := uc.userService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to get user",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    user,
		"status":  true,
	})
}

func (uc *UserController) FindByID(c *gin.Context) {
	userID := c.Param("id")
	user, err := uc.userService.FindByID(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to get user",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    user,
		"status":  true,
	})
}

func (uc *UserController) Delete(c *gin.Context) {
	userID := c.Param("id")
	user, err := uc.userService.Delete(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to Delete The User",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Delete User is Successfully",
		"data":    user,
	})
}
