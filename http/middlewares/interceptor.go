package middlewares

import (
	"fmt"
	"frankie060392/rest-api-clean-arch/bootstrap"
	"frankie060392/rest-api-clean-arch/http/common"
	"frankie060392/rest-api-clean-arch/internal/user/model"
	cryptography "frankie060392/rest-api-clean-arch/pkg/crypto"

	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func DeserializeUser(us model.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader("Authorization")
		fields := strings.Fields(authorizationHeader)
		var accessToken string
		if len(fields) != 0 && fields[0] == "Bearer" {
			accessToken = fields[1]
		}

		if accessToken == "" {
			ctx.JSON(http.StatusUnauthorized, common.ResponseData{Status: false, Message: http.StatusText(http.StatusUnauthorized)})
			return
		}

		config := bootstrap.LoadConfig(".")

		sub, err := cryptography.ValidateToken(accessToken, config.AccessTokenPublicKey)

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, common.ResponseData{Status: false, Message: http.StatusText(http.StatusUnauthorized)})
			return
		}

		user, err := us.GetByID(ctx, fmt.Sprint(sub))

		if err != nil {
			ctx.JSON(http.StatusNotFound, common.ResponseData{Status: false, Message: http.StatusText(http.StatusNotFound)})
		}

		ctx.Set("currentUser", user)
		ctx.Next()
	}
}
