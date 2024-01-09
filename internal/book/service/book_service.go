package service

import (
	"context"
	"frankie060392/rest-api-clean-arch/internal/book/model"
)

type bookService struct {
	repo model.BookRepositoryInterface
}

// Create implements model.BookServiceInterface.
func (*bookService) Create(c context.Context, book *model.Book) error {
	panic("unimplemented")
}

// GetById implements model.BookServiceInterface.
func (*bookService) GetById(c context.Context, id string) (model.Book, error) {
	panic("unimplemented")
}

func NewBookService(repo model.BookRepositoryInterface) model.BookServiceInterface {
	return &bookService{repo: repo}
}
