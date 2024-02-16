package initializers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	if os.Getenv("GITHUB_ACTIONS") != "true" {
		// Load from .env file for local development
		err := godotenv.Load()

		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}