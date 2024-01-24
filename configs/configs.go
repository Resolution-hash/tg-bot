package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TelegramAPIToken, DbHost, DbPort, DbUser, DbPass, DbName string
}

func LoadConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("failed load file .env")
	}

	return &Config{
		TelegramAPIToken: os.Getenv("TELEGRAM_API_TOKEN"),
		DbHost: os.Getenv("DB_HOST"),
		DbPort: os.Getenv("DB_PORT"),
		DbUser: os.Getenv("DB_USER"),
		DbPass: os.Getenv("DB_PASS"),
		DbName: os.Getenv("DB_NAME"),
	}
}
