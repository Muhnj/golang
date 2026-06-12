package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	BaseURL       string
	ConsumerKey   string
	ConsumerSecret string
	MerchantCode  string
	ServerPort    string
}

func Load() Config {
	_ = godotenv.Load()

	return Config{
		BaseURL:        os.Getenv("PAYLEO_BASE_URL"),
		ConsumerKey:    os.Getenv("PAYLEO_CONSUMER_KEY"),
		ConsumerSecret: os.Getenv("PAYLEO_CONSUMER_SECRET"),
		MerchantCode:   os.Getenv("PAYLEO_MERCHANT_CODE"),
		ServerPort:     os.Getenv("SERVER_PORT"),
	}
}