package main

import (
	"newsbot/configs"
  "newsbot/internal/telegram/bot"
)

func main() {
  cfg := configs.LoadConfig()
	bot.StartBot(cfg.TelegramAPIToken)
}
