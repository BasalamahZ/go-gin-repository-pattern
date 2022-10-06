package services

import (
	"go-gorm-jwt/models"
	"go-gorm-jwt/repositories"
)

type BookService interface {
	FindAll() ([]models.Book, error)
	FindByID(bookID string) (models.Book, error)
	Create(book models.Book) (models.Book, error)
	Update(book models.Book, bookID string) (models.Book, error)
	Delete(bookID string) (models.Book, error)
}

type bookService struct {
	bookRepository repositories.BookRepository
}

// Create implements BookService
func (bs *bookService) Create(book models.Book) (models.Book, error) {
	book, err := bs.bookRepository.Create(book)
	if err != nil {
		return book, err
	}
	return book, nil
}

// FindAll implements BookService
func (bs *bookService) FindAll() ([]models.Book, error) {
	book, err := bs.bookRepository.FindAll()
	if err != nil {
		return book, err
	}
	return book, nil
}

// FindByID implements BookService
func (bs *bookService) FindByID(bookID string) (models.Book, error) {
	book, err := bs.bookRepository.FindByID(bookID)
	if err != nil {
		return book, err
	}
	return book, nil
}

// Update implements BookService
func (bs *bookService) Update(book models.Book, bookID string) (models.Book, error) {
	newBook := models.Book{
		Title:       book.Title,
		Description: book.Description,
		Price:       book.Price,
		Author:      book.Author,
		Rating:      book.Rating,
	}

	book, err := bs.bookRepository.Update(newBook, bookID)
	if err != nil {
		return book, err
	}
	return book, nil
}

// Delete implements BookService
func (bs *bookService) Delete(bookID string) (models.Book, error) {
	book, err := bs.bookRepository.Delete(bookID)
	if err != nil {
		return book, err
	}
	return book, nil
}

func NewBookService(bookRepository *repositories.BookRepository) BookService {
	return &bookService{
		bookRepository: *bookRepository,
	}
}
