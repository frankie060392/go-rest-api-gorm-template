package handlers

import (
	"frankie060392/rest-api-clean-arch/http/common"
	"frankie060392/rest-api-clean-arch/internal/user/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService model.UserServiceInterface
}

func NewUserHandler(userService model.UserServiceInterface) userHandler {
	return userHandler{
		userService: userService,
	}
}

func (uh *userHandler) GetById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, nil)
}

func (uh *userHandler) GetUser(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(model.User)
	userResponse := &model.UserResponse{
		ID:        currentUser.ID,
		Name:      currentUser.Name,
		Email:     currentUser.Email,
		CreatedAt: currentUser.CreatedAt,
		UpdatedAt: currentUser.UpdatedAt,
	}
	ctx.JSON(http.StatusOK, common.ResponseData{Status: true, Message: "", Data: userResponse})
}
