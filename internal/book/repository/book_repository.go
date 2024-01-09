package repository

import (
	"context"
	"frankie060392/rest-api-clean-arch/internal/book/model"

	"gorm.io/gorm"
)

type bookRepository struct {
	DB *gorm.DB
}

// Create implements model.BookRepositoryInterface.
func (r *bookRepository) Create(c context.Context, book *model.Book) error {
	result := r.DB.Save(&book)
	return result.Error
}

// GetById implements model.BookRepositoryInterface.
func (r *bookRepository) GetById(c context.Context, id string) (model.Book, error) {
	var book model.Book
	result := r.DB.First(&book, "id = ?", id)
	return book, result.Error
}

func NewBookRepository(db *gorm.DB) model.BookRepositoryInterface {
	return &bookRepository{DB: db}
}
