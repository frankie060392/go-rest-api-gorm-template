package repository

import (
	"context"
	"frankie060392/rest-api-clean-arch/internal/user/model"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
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

func NewRepository(db *gorm.DB) model.UserRepository {
	return &repository{DB: db}
}
