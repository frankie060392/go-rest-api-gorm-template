package handlers

import (
	"frankie060392/rest-api-clean-arch/bootstrap"
	"frankie060392/rest-api-clean-arch/http/common"
	"frankie060392/rest-api-clean-arch/http/messages"
	"frankie060392/rest-api-clean-arch/internal/user/model"
	cryptography "frankie060392/rest-api-clean-arch/pkg/crypto"
	"frankie060392/rest-api-clean-arch/pkg/utils"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type authHandler struct {
	userService model.UserServiceInterface
}

func NewAuthHandler(userService model.UserServiceInterface) authHandler {
	return authHandler{
		userService: userService,
	}
}

func (ah *authHandler) SignUp(ctx *gin.Context) {
	var payload *model.SignUpInput
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, common.ResponseData{Status: false, Message: messages.ErrorCreate})
		return
	}
	if payload.Password != payload.PasswordConfirm {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.ResponseData{Status: false, Message: messages.ErrorCreate})
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
	err = ah.userService.CreateUser(ctx, &newUser)
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
	ctx.JSON(http.StatusCreated, common.ResponseData{Status: true, Data: userResponse, Message: messages.SuccessCreate})
}

func (auth *authHandler) SignIn(ctx *gin.Context) {
	var payload *model.SignInInput

	err := ctx.ShouldBindJSON(&payload)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.ResponseData{Status: false, Message: http.StatusText(http.StatusBadRequest)})
		return
	}

	user, err := auth.userService.GetByEmail(ctx, payload.Email)
	if err != nil {
		ctx.JSON(http.StatusNotFound, common.ResponseData{Status: false, Message: http.StatusText(http.StatusNotFound)})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.ResponseData{Status: false, Message: http.StatusText(http.StatusBadRequest)})
		return
	}
	err = utils.VerifyPassword(user.Password, payload.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.ResponseData{Status: false, Message: messages.ErrorUnauthorized})
		return
	}

	config := bootstrap.LoadConfig(".")
	accessToken, err := cryptography.CreateToken(config.AccessTokenExpiresIn, user.Email, config.AccessTokenPrivateKey)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.ResponseData{Status: false, Message: messages.ErrorUnauthorized})
		return
	}

	ctx.JSON(http.StatusOK, common.ResponseData{Status: true, Message: messages.SuccessLogin, Data: accessToken})
}
