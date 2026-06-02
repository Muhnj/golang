package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort string

	APIURL string
	User   string
	Pass   string
}

func Load() *Config {

	err := godotenv.Load()

	if err != nil {
		log.Println(".env file not found")
	}

	return &Config{
		AppPort: os.Getenv("APP_PORT"),

		APIURL: os.Getenv("API_URL"),
		User:   os.Getenv("API_USERNAME"),
		Pass:   os.Getenv("API_PASSWORD"),
	}
}
