package main

import (
	"frankie060392/rest-api-clean-arch/bootstrap"
	"frankie060392/rest-api-clean-arch/http/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	config := bootstrap.LoadConfig(".")
	DB := bootstrap.NewDbConnection(config)

	gin := gin.Default()

	routes.Setup(DB, gin)
	gin.Run(config.SERVER_ADDRESS)
}
