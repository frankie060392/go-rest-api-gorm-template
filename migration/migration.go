package main

import (
	"fmt"
	"frankie060392/rest-api-clean-arch/bootstrap"
	"frankie060392/rest-api-clean-arch/internal/user/model"
)

func main() {
	config := bootstrap.LoadConfig(".")
	DB := bootstrap.NewDbConnection(config)
	DB.AutoMigrate(&model.User{})
	fmt.Println("? Migration complete")
}
