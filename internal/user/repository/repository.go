package repository

import (
	"context"
	"frankie060392/rest-api-clean-arch/internal/user/model"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

// GetByEmail implements model.UserRepositoryInterface.
func (r *repository) GetByEmail(c context.Context, email string) (model.User, error) {
	var user model.User
	result := r.DB.First(&user, "email = ?", email)
	return user, result.Error
}

// CreateUser implements model.UserRepository.
func (r *repository) CreateUser(c context.Context, user *model.User) error {
	result := r.DB.Save(&user)
	return result.Error
}

// GetByID implements model.UserRepository.
func (r *repository) GetByID(c context.Context, id string) (model.User, error) {
	var user model.User
	result := r.DB.First(&user, "id = ?", id)
	return user, result.Error
}

func NewUserRepository(db *gorm.DB) model.UserRepositoryInterface {
	return &repository{DB: db}
}
