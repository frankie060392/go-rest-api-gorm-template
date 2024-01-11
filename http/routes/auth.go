package routes

import (
	"frankie060392/rest-api-clean-arch/http/handlers"
	"frankie060392/rest-api-clean-arch/internal/user/repository"
	"frankie060392/rest-api-clean-arch/internal/user/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewAuthRouter(DB *gorm.DB, group *gin.RouterGroup) {

	router := group.Group("auth")

	ur := repository.NewUserRepository(DB)
	us := service.NewUserService(ur)
	ah := handlers.NewAuthHandler(us)
	router.POST("/signin", ah.SignIn)
	router.POST("/signup", ah.SignUp)
}
