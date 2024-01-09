package service

import (
	"context"
	"frankie060392/rest-api-clean-arch/internal/user/model"
)

type service struct {
	repo model.UserRepository
}

// CreateUser implements model.UserService.
func (s *service) CreateUser(c context.Context, user model.User) (model.User, error) {
	return s.repo.CreateUser(c, user)
}

// GetByID implements model.UserService.
func (s *service) GetByID(c context.Context, id string) (model.User, error) {
	return s.repo.GetByID(c, id)
}

func NewService(r model.UserRepository) model.UserService {
	return &service{
		repo: r,
	}
}
