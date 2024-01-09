package handlers

import (
	"frankie060392/rest-api-clean-arch/internal/user/model"
	"frankie060392/rest-api-clean-arch/pkg/utils"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService model.UserService
}

func NewUserHandler(userService model.UserService) UserHandler {
	return UserHandler{
		userService: userService,
	}
}

func (uh *UserHandler) Create(ctx *gin.Context) {
	var payload *model.SignUpInput
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	if payload.Password != payload.PasswordConfirm {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	now := time.Now()
	newUser := model.User{
		Name:      payload.Name,
		Email:     strings.ToLower(payload.Email),
		Password:  hashedPassword,
		CreatedAt: now,
		UpdatedAt: now,
	}
	err = uh.userService.CreateUser(ctx, &newUser)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, nil)
		return
	}

	userResponse := &model.UserResponse{
		ID:        newUser.ID,
		Name:      newUser.Name,
		Email:     newUser.Email,
		CreatedAt: newUser.CreatedAt,
		UpdatedAt: newUser.UpdatedAt,
	}
	ctx.JSON(http.StatusCreated, userResponse)
}

func (uh *UserHandler) GetById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, nil)
}
