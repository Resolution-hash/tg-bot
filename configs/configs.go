package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TelegramAPIToken string
	DatabaseURL      string
}

func LoadConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("failed load file .env")
	}

	token := os.Getenv("TELEGRAM_API_TOKEN")
	databaseUrl := os.Getenv("DATABASE_URL")

	return &Config{
		TelegramAPIToken: token,
		DatabaseURL:      databaseUrl,
	}
}
