package routes

import (
	"frankie060392/rest-api-clean-arch/http/handlers"
	"frankie060392/rest-api-clean-arch/internal/book/repository"
	"frankie060392/rest-api-clean-arch/internal/book/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewBookRouter(db *gorm.DB, group *gin.RouterGroup) {
	router := group.Group("book")
	ur := repository.NewBookRepository(db)
	us := service.NewBookService(ur)
	uh := handlers.NewBookHandler(us)
	router.POST("/", uh.Create)
	router.GET("/:id", uh.GetByID)
}
