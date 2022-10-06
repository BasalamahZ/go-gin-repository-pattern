package controllers

import (
	"go-gorm-jwt/models"
	"go-gorm-jwt/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	bookService services.BookService
}

func NewBookController(bookService *services.BookService) BookController {
	return BookController{
		bookService: *bookService,
	}
}

func (bc *BookController) Create(c *gin.Context) {
	bookRequest := new(models.Book)
	if c.Bind(&bookRequest) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to read body",
		})
		return
	}
	books, err := bc.bookService.Create(*bookRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to Create Book",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "success",
		"data":    books,
	})
}

func (bc *BookController) FindAll(c *gin.Context) {
	books, err := bc.bookService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to get book",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    books,
		"status":  true,
	})
}

func (bc *BookController) FindByID(c *gin.Context) {
	bookID := c.Param("id")
	books, err := bc.bookService.FindByID(bookID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to get book",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    books,
		"status":  true,
	})
}

func (bc *BookController) Update(c *gin.Context) {
	bookID := c.Param("id")
	bookRequest := new(models.Book)
	c.BindJSON(&bookRequest)
	book, err := bc.bookService.Update(*bookRequest, bookID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to Update Book",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "success",
		"data":    book,
	})
}

func (bc *BookController) Delete(c *gin.Context) {
	bookID := c.Param("id")

	book, err := bc.bookService.Delete(bookID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to Delete The Book",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"status":  true,
		"data":    book,
	})
}
