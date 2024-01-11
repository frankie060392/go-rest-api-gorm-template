package routes

import (
	"frankie060392/rest-api-clean-arch/http/handlers"
	"frankie060392/rest-api-clean-arch/internal/user/repository"
	"frankie060392/rest-api-clean-arch/internal/user/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewUserRouter(db *gorm.DB, group *gin.RouterGroup) {
	router := group.Group("user")
	ur := repository.NewUserRepository(db)
	us := service.NewUserService(ur)
	uh := handlers.NewUserHandler(us)
	router.GET("/", uh.GetUser)
}
