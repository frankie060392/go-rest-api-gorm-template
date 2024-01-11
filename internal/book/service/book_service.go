package service

import (
	"context"
	"frankie060392/rest-api-clean-arch/internal/book/model"
)

type bookService struct {
	repo model.BookRepositoryInterface
}

// GetBooksByAuthor implements model.BookServiceInterface.
func (s *bookService) GetBooksByAuthor(c context.Context, email string) ([]model.Book, error) {
	return s.repo.GetBooksByAuthor(c, email)
}

// Create implements model.BookServiceInterface.
func (s *bookService) Create(c context.Context, book *model.Book) error {
	return s.repo.Create(c, book)
}

// GetById implements model.BookServiceInterface.
func (s *bookService) GetById(c context.Context, id string) (model.Book, error) {
	return s.repo.GetById(c, id)
}

func NewBookService(repo model.BookRepositoryInterface) model.BookServiceInterface {
	return &bookService{repo: repo}
}
