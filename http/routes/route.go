package routes

import (
	"frankie060392/rest-api-clean-arch/http/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB, gin *gin.Engine) {
	// Public APIS
	publicRouter := gin.Group("api")
	NewAuthRouter(db, publicRouter)

	// Private APIs
	protectedRouter := gin.Group("api")
	protectedRouter.Use(middlewares.DeserializeUser(db))
	NewUserRouter(db, protectedRouter)
	NewBookRouter(db, protectedRouter)
}
