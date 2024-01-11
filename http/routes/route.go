package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB, gin *gin.Engine) {
	protectedRouter := gin.Group("api")
	// Middleware to verify AccessToken
	// All Private APIs
	NewUserRouter(db, protectedRouter)
	NewBookRouter(db, protectedRouter)
	NewAuthRouter(db, protectedRouter)
}
