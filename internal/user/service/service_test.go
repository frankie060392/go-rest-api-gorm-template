package service_test

import (
	"context"
	"frankie060392/rest-api-clean-arch/internal/mocks"
	"frankie060392/rest-api-clean-arch/internal/user/model"
	"frankie060392/rest-api-clean-arch/internal/user/service"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFetchByUserID(t *testing.T) {
	mockUserRepository := new(mocks.UserRepository)
	userObjectID := uuid.New()
	userID := userObjectID.String()

	t.Run("success", func(t *testing.T) {

		mockTask := model.User{
			ID:       uuid.New(),
			Name:     "frankie",
			Email:    "frankie060392@gmail.com",
			Password: "frankietest",
		}

		mockUserRepository.On("GetByID", mock.Anything, userID).Return(mockTask, nil)

		u := service.NewUserService(mockUserRepository)
		createdUser, err := u.GetByID(context.Background(), userID)
		assert.NoError(t, err)
		assert.NotNil(t, createdUser)

		mockUserRepository.AssertExpectations(t)
	})

}
