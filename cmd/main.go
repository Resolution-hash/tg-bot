package main

import (
  "log"
	"newsbot/configs"
	"newsbot/internal/telegram/bot"
)

func main() {
  cfg := configs.LoadConfig()
  StartTelegramBot(cfg.TelegramAPIToken)
}


func StartTelegramBot(TelegramAPIToken string){
  err := bot.StartBot(TelegramAPIToken)
  if err != nil{
    log.Fatalf("The bot didn`t start. Error: %s", err)
  }
}
