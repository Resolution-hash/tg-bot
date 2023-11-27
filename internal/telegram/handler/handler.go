// package handler

// import (
// 	"log"

// 	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
// )

// func HandleUpdates(updates tgbotapi.UpdatesChannel, bot *tgbotapi.BotAPI) {
// 	for update := range updates {
// 		if update.Message == nil {
// 			continue
// 		}
// 		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

// 		sender.SendArticle(update.Message.Chat.ID, update.Message.Text, bot)
// 	}
// }
