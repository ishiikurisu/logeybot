package main

import (
	telegram "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ishiikurisu/logeybot/controller"
	"os"
)

// The entry point of this application. Creates a Telegram bot
func main() {
	bot, err := telegram.NewBotAPI(os.Args[1])
	if err != nil {
		panic(err)
	}
	bot.Debug = true
	u := telegram.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	controllers := make(map[int64]*controller.Controller)
	for update := range updates {
		if update.Message != nil {
			identification := update.Message.Chat.ID
			if _, ok := controllers[identification]; !ok {
				dummy := controller.NewController(identification)
				controllers[identification] = &dummy
			}
			// TODO Adapt this to receive files as well
			var msg telegram.Chattable
			if controller.GetMessageKind(update.Message.Text) == "text" {
				msg = telegram.NewMessage(identification, controllers[identification].Listen(update.Message.Text))
			} else {
				msg = telegram.NewDocumentUpload(identification, controllers[identification].Listen(update.Message.Text))
			}
			bot.Send(msg)
		}
	}
}
