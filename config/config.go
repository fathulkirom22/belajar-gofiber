package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func InitConfig() {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
}

// Config func to get env value
func Config(key string) string {
	return os.Getenv(key)
}
