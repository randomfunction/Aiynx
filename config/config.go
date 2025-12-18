package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	Port        string
}

func LoadConfig() *Config {
	godotenv.Load()
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is required")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // Default port
	}

	return &Config{
		DatabaseURL: dbURL,
		Port:        port,
	}
}
