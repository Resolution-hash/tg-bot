package bot

import (
	"log"
	"newsbot/internal/telegram/sender"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func StartBot(token string) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true
	log.Println("Bot started!")

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		sender.SendArticle(update.Message.Chat.ID, bot)
	}
}

func setCommand(){
	
}
