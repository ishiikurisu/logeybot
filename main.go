package main

import (
    "os"
	telegram "github.com/go-telegram-bot-api/telegram-bot-api"
    "github.com/ishiikurisu/logeybot/controller"
)

// The entry point of this application. Creates a Telegram bot
func main() {
    // Setup bot
	bot, err := telegram.NewBotAPI(os.Args[1])
	if err != nil {
        panic(err)
	}
	bot.Debug = true
	u := telegram.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
    // Load previous states from memory
    controllers := make(map[int64]*controller.Controller)
    // Putting bot to work depending on messages
	for update := range updates {
		if update.Message != nil {
            // TODO Do this on parallel
            identification := update.Message.Chat.ID
            if _, ok := controllers[identification]; !ok {
                dummy := controller.NewController(identification)
                controllers[identification] = &dummy
            }
    		msg := telegram.NewMessage(identification, controllers[identification].Listen(update.Message.Text))
    		bot.Send(msg)
		}
	}
}
