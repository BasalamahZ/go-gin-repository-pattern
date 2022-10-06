package main

import (
	"go-gorm-jwt/configs"
	"go-gorm-jwt/controllers"
	"go-gorm-jwt/middleware"
	"go-gorm-jwt/repositories"
	"go-gorm-jwt/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	configs.LoadEnvVariables()
	configs.ConnectDB()
	configs.SyncDB()
}

func main() {
	server := gin.Default()

	server.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "home",
		})
	})

	authRepository := repositories.NewAuthRepository()
	authService := services.NewAuthService(&authRepository)
	authController := controllers.NewAuthController(&authService)
	server.POST("/signup", authController.Register)
	server.POST("/login", authController.Login)

	userRepository := repositories.NewUserRepository()
	userService := services.NewUserService(&userRepository)
	userController := controllers.NewUserController(&userService)
	server.GET("/user", middleware.VerifyAuth, userController.FindAll)
	server.GET("/user/:id", middleware.VerifyAuth, userController.FindByID)
	server.DELETE("/user/:id", middleware.VerifyAuth, userController.Delete)

	bookRepository := repositories.NewBookRepository()
	bookService := services.NewBookService(&bookRepository)
	bookController := controllers.NewBookController(&bookService)
	server.POST("/book", middleware.VerifyAuth, bookController.Create)
	server.GET("/book", middleware.VerifyAuth, bookController.FindAll)
	server.GET("/book/:id", middleware.VerifyAuth, bookController.FindByID)
	server.PUT("/book/:id", middleware.VerifyAuth, bookController.Update)
	server.DELETE("/book/:id", middleware.VerifyAuth, bookController.Delete)

	server.Run()
}
