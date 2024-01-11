package service

import (
	"context"
	"frankie060392/rest-api-clean-arch/internal/user/model"
)

type userService struct {
	repo model.UserRepositoryInterface
}

// GetByEmail implements model.UserServiceInterface.
func (s *userService) GetByEmail(c context.Context, email string) (model.User, error) {
	return s.repo.GetByEmail(c, email)
}

// CreateUser implements model.UserService.
func (s *userService) CreateUser(c context.Context, user *model.User) error {
	return s.repo.CreateUser(c, user)
}

// GetByID implements model.UserService.
func (s *userService) GetByID(c context.Context, id string) (model.User, error) {
	return s.repo.GetByID(c, id)
}

func NewUserService(r model.UserRepositoryInterface) model.UserServiceInterface {
	return &userService{
		repo: r,
	}
}
