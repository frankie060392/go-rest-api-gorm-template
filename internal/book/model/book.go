package model

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name      string    `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type BookCreate struct {
	Name      string    `json:"name" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type BookRepositoryInterface interface {
	Create(c context.Context, book *Book) error
	GetById(c context.Context, id string) (Book, error)
}

type BookServiceInterface interface {
	Create(c context.Context, book *Book) error
	GetById(c context.Context, id string) (Book, error)
}
