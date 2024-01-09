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
func (*bookRepository) Create(c context.Context, book *model.Book) error {
	panic("unimplemented")
}

// GetById implements model.BookRepositoryInterface.
func (*bookRepository) GetById(c context.Context, id string) (model.Book, error) {
	panic("unimplemented")
}

func NewBookRepository(db *gorm.DB) model.BookRepositoryInterface {
	return &bookRepository{DB: db}
}
