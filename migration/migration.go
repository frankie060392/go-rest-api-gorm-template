package migration

import (
	"fmt"
	"frankie060392/rest-api-clean-arch/bootstrap"
	"frankie060392/rest-api-clean-arch/internal/user/model"
	"log"
)

func init() {
	config, err := bootstrap.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	bootstrap.ConnectDB(&config)
}

func main() {
	bootstrap.DB.AutoMigrate(&model.User{})
	fmt.Println("? Migration complete")
}
