package main

import "fmt"
import telegram "github.com/go-telegram-bot-api/telegram-bot-api"

// Code from https://github.com/go-telegram-bot-api/telegram-bot-api and
// adapted to run on this application
func main() {
    fmt.Printf("Let's do this!\n")

    // Registering bot
    var token string
    fmt.Scanln(&token)
    bot, err := telegram.NewBotAPI(token)
    if err != nil {
        fmt.Println(err)
        panic(err)
    }
    bot.Debug = true
    fmt.Printf("Authorized on account %s\n", bot.Self.UserName)

    // Getting updates
    u := telegram.NewUpdate(0)
    u.Timeout = 60
    updates, err := bot.GetUpdatesChan(u)
    for update := range updates {
        if update.Message == nil {
            continue
        }

        fmt.Printf("[%s] %s\n", update.Message.From.UserName, update.Message.Text)
        msg := telegram.NewMessage(update.Message.Chat.ID, update.Message.Text)
        msg.ReplyToMessageID = update.Message.MessageID

        bot.Send(msg)
    }
}
