package config

import (
	"log"
	"os"
)

type Config struct{
	DBHost string
	DBPort string
	DBUser string
	DBPassword string
	DBName string
	HTTPAddr string
}

func Load() Config {
	cfg := Config{
		DBHost: mustGet("DB_HOST"),
		DBPort: mustGet("DB_PORT"),
		DBUser: mustGet("DB_USER"),
		DBPassword: mustGet("DB_PASSWORD"),
		DBName: mustGet("DB_NAME"),
		HTTPAddr: mustGet("HTTP_ADDR"),
	}

	return cfg
}

func mustGet(key string) (val string) {
	val = os.Getenv(key)
	if val == "" {
		log.Fatalf("Missing required environment variable: %s", key)
	}

	return
}