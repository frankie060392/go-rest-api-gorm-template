package handlers

import (
	"frankie060392/rest-api-clean-arch/internal/user/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService model.UserServiceInterface
}

func NewUserHandler(userService model.UserServiceInterface) UserHandler {
	return UserHandler{
		userService: userService,
	}
}

func (uh *UserHandler) GetById(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, nil)
}
