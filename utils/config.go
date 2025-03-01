package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MONGODB_URI     string
	DB_NAME         string
	AWANLLM_API_KEY string
}

var AppConfig *Config

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading evn file")
	}

	AppConfig = &Config{
		MONGODB_URI:     os.Getenv("MONGODB_URI"),
		DB_NAME:         os.Getenv("DB_NAME"),
		AWANLLM_API_KEY: os.Getenv("AWANLLM_API_KEY"),
	}
}
