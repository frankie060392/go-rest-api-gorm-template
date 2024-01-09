package main

import (
	"fmt"
	"frankie060392/rest-api-clean-arch/bootstrap"
	bookModel "frankie060392/rest-api-clean-arch/internal/book/model"
	userModel "frankie060392/rest-api-clean-arch/internal/user/model"
)

func main() {
	config := bootstrap.LoadConfig(".")
	DB := bootstrap.NewDbConnection(config)
	DB.AutoMigrate(&bookModel.Book{})
	DB.AutoMigrate(&userModel.User{})
	fmt.Println("? Migration complete")
}
