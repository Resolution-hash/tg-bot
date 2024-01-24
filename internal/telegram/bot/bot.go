package bot

import (
	"log"
	"newsbot/internal/models"
	postmanager "newsbot/internal/postManager"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const HelloMsg = "Привет! Я показаю новости о языке Golang с сайта habr!\nВыбери команду, чтобы я показал новости."
const DefaultMsg = "Я не знаю такой команды ):"

func StartBot(token string) error {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
		return err
	}

	bot.Debug = true
	log.Println("Bot launched!")
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatal(err)
		return err
	}

	for update := range updates {
		if update.CallbackQuery != nil {
			handleCallback(update, bot)
		}
		if update.Message != nil {
			log.Println("\n\n\n" + update.Message.Text)
			handleCommand(update, bot)
		}

	}
	return nil
}

func handleCommand(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	switch update.Message.Command() {
	case "start":
		sendMessage(update, bot, HelloMsg)
	// case "add":
	// 	parser.InsertData()
	// case "del":
	// 	parser.DeleteAllRows()
	// case "downloadall":
	// 	parser.AllArticles()
	// case "random":
	// 	post := postmanager.RandomPost()
	// 	sendPost(post, update, bot)
	default:
		sendMessage(update, bot, DefaultMsg)
	}
}

func handleCallback(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	callback := update.CallbackQuery
	data := callback.Data
	log.Println(data)
	switch data {
	case "random":
		post := postmanager.RandomPost()
		sendPost(post, update, bot)
	case "today":
		posts := postmanager.TodayPosts()
		sendPost(posts, update, bot)
	default:
		sendMessage(update, bot, DefaultMsg)
	}

}

func sendMessage(update tgbotapi.Update, bot *tgbotapi.BotAPI, message string) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Random", "random"),
			tgbotapi.NewInlineKeyboardButtonData("Today", "today"),
		),
		// tgbotapi.NewInlineKeyboardRow(
		// 	tgbotapi.NewInlineKeyboardButtonData("Last week", "lstweek"),
		// ),
	)

	msg.ReplyMarkup = &keyboard
	_, err := bot.Send(msg)
	if err != nil {
		log.Println("Error send message")
	}
}

func sendPost(posts []models.Article, update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	callback := update.CallbackQuery
	if len(posts) == 0 || (len(posts) == 1 && posts[0].ID == 0) {
		text := "Сегодня нет постов.\nВыбери другую комнду"
		msg := tgbotapi.NewMessage(callback.Message.Chat.ID, text)
		keyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Random", "random"),
				tgbotapi.NewInlineKeyboardButtonData("Today", "today"),
			),
			// tgbotapi.NewInlineKeyboardRow(
			// 	tgbotapi.NewInlineKeyboardButtonData("Last week", "lastweek"),
			// ),
		)
		msg.ReplyMarkup = &keyboard

		log.Println("post is send!")
		log.Println(posts)
		bot.Send(msg)
		return
	}
	for _, post := range posts {
		log.Println("sendPost starts!")
		text := "Дата публикации: " + post.Date + "\n\n" + post.Title + "\n\n" + post.Content + "\n\n" + post.Url + "\n"
		msg := tgbotapi.NewMessage(callback.Message.Chat.ID, text)

		keyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Random", "random"),
				tgbotapi.NewInlineKeyboardButtonData("Today", "today"),
			),
			// tgbotapi.NewInlineKeyboardRow(
			// 	tgbotapi.NewInlineKeyboardButtonData("Last week", "lstweek"),
			// ),
		)
		msg.ReplyMarkup = &keyboard

		log.Println("post is send!")
		log.Println(posts)
		bot.Send(msg)
	}
}
