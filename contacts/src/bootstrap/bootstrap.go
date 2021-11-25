package bootstrap

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/paulolancao/go-contacts/di"
	"github.com/paulolancao/go-contacts/logging"
	"github.com/paulolancao/go-contacts/repository"
	"github.com/paulolancao/go-contacts/services"
	"github.com/paulolancao/go-contacts/webserver"
)

// Init Bootstrap
func Init() {
	// Load .env file
	loadDotEnv()

	// Logging startup
	logging.Init()

	// Repositories
	baserepository := repository.Init()

	// Services
	service := services.Init(&baserepository)

	// Wrapper of repositories or and services
	di := di.DI{ContactRepository: &baserepository, ContactService: &service}

	// Webserver
	webserver.Init(&di)
}

func loadDotEnv() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Print(err)
	}
}
