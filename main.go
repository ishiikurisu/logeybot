package main

import (
    "os"
	"github.com/go-telegram-bot-api/telegram-bot-api"
    "github.com/ishiikurisu/logeybot/controller"
)

func main() {
    // Setup bot
	bot, err := tgbotapi.NewBotAPI(os.Args[1])
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
    // Setup memory
    controllers := make(map[int64]*controller.Controller)
    // Putting bot to work depending on messages
	for update := range updates {
		if update.Message == nil {
			continue
		}

        identification := update.Message.Chat.ID
        _, ok := controllers[identification]
        if !ok {
            e := controller.NewController(identification)
            controllers[identification] = &e
        }
        c := controllers[identification]
		msg := tgbotapi.NewMessage(identification, c.Listen(update.Message.Text))
		bot.Send(msg)
	}
}
