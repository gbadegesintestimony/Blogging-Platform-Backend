package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	// Implementation to load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func GetEnv(key string) string {
	// Implementation to get environment variable by key
	return os.Getenv(key)
}
