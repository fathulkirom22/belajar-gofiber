package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// InitConfig : init config
func InitConfig() {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
}

// GetConfig : Config func to get env value
func GetConfig(key string) string {
	return os.Getenv(key)
}
