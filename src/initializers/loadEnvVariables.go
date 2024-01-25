// create a function for loading environment variables
package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

// the name of the function should be in capital letter so that to load in other repository
func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
