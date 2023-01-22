package helpers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key, fallback string) string {

	// load .env file
	err := godotenv.Load()
  
	if err != nil {
	  log.Fatalf("Error loading .env file")
	}

	value := os.Getenv(key)

	if len(value) == 0 {
        return fallback
    }
  
	return value
}