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
	br := repository.NewBookRepository(db)
	bs := service.NewBookService(br)
	bh := handlers.NewBookHandler(bs)
	router.POST("/", bh.Create)
	router.GET("/:id", bh.GetByID)
	router.GET("/list", bh.GetBooksByUser)
}
