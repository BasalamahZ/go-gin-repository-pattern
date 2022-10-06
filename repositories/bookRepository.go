package repositories

import (
	"go-gorm-jwt/configs"
	"go-gorm-jwt/models"
)

type BookRepository interface {
	FindAll() ([]models.Book, error)
	FindByID(bookID string) (models.Book, error)
	Create(book models.Book) (models.Book, error)
	Update(book models.Book, bookID string) (models.Book, error)
	Delete(bookID string) (models.Book, error)
}

type bookRepository struct {
}

// Create implements BookRepository
func (*bookRepository) Create(book models.Book) (models.Book, error) {
	db := configs.DB.Create(&book)
	if db != nil {
		return book, db.Error
	}
	return book, nil
}

// FindAll implements BookRepository
func (*bookRepository) FindAll() ([]models.Book, error) {
	var book []models.Book
	db := configs.DB.Find(&book)
	if db != nil {
		return book, db.Error
	}
	return book, nil
}

// FindByID implements BookRepository
func (*bookRepository) FindByID(bookID string) (models.Book, error) {
	var book models.Book
	db := configs.DB.First(&book, bookID)
	if db != nil {
		return book, db.Error
	}
	return book, nil
}

// Update implements BookRepository
func (*bookRepository) Update(book models.Book, bookID string) (models.Book, error) {
	db := configs.DB.Model(&book).Where("id = ?", bookID).Updates(&book)
	if db != nil {
		return book, db.Error
	}
	return book, nil
}

// Delete implements BookRepository
func (*bookRepository) Delete(bookID string) (models.Book, error) {
	var book models.Book
	db := configs.DB.Delete(bookID)
	if db != nil {
		return book, db.Error
	}
	return book, nil
}

func NewBookRepository() BookRepository {
	return &bookRepository{}
}
