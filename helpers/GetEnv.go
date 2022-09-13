package helpers

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetEnv(key, static string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return static
}
