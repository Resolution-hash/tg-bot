package sender

import (
	"fmt"
	"log"
	"newsbot/internal/parser"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func SendArticle(chatID int64, bot *tgbotapi.BotAPI) {
	articles := parser.GetArticles()
	post := getArticleText(articles)
	msg := tgbotapi.NewMessage(chatID, post)
	_, err := bot.Send(msg)
	if err != nil {
		log.Println(err)
	}
}

// Url, Image, Name, Content, DateTime string
func getArticleText(articles []parser.Article) string {
	fmt.Println(articles[1])
	text := "Ссылка:" + articles[1].Url +
		"\nКартинка:" + articles[1].Image +
		"\nНазвание статьи:" + articles[1].Name +
		"\nСодержание:" + articles[1].Content +
		"\nДата публикации:" + articles[1].DateTime
	return text
}
