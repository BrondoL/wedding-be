package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	APP_HOST    string
	APP_PORT    string
	API_VERSION string
	APP_URL     string
	APP_ENV     string
	CLIENT_URL  string
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASS     string
	DB_NAME     string
	DB_SSL_MODE string
	REDIS_HOST  string
	REDIS_PORT  string
	REDIS_PASS  string
}

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
	}
}

func GetEnv() Config {
	cfg := Config{
		APP_HOST:    os.Getenv("APP_HOST"),
		APP_PORT:    os.Getenv("APP_PORT"),
		API_VERSION: os.Getenv("API_VERSION"),
		APP_URL:     os.Getenv("APP_URL"),
		APP_ENV:     os.Getenv("APP_ENV"),
		CLIENT_URL:  os.Getenv("CLIENT_URL"),
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_USER:     os.Getenv("DB_USER"),
		DB_PASS:     os.Getenv("DB_PASS"),
		DB_NAME:     os.Getenv("DB_NAME"),
		DB_SSL_MODE: os.Getenv("DB_SSL_MODE"),
		REDIS_HOST:  os.Getenv("REDIS_HOST"),
		REDIS_PORT:  os.Getenv("REDIS_PORT"),
		REDIS_PASS:  os.Getenv("REDIS_PASS"),
	}

	return cfg
}
