package routes

import (
	"frankie060392/rest-api-clean-arch/http/handlers"

	"github.com/gin-gonic/gin"
)

type UserRoute struct {
	userHandler handlers.UserHandler
}

func NewUserRoute(userHandler handlers.UserHandler) UserRoute {
	return UserRoute{userHandler}
}

func (ur *UserRoute) Register(rg *gin.RouterGroup) {

	router := rg.Group("users")
	router.GET("/", ur.userHandler.GetById)
	router.POST("/", ur.userHandler.Create)
}
