package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"iam43x.vpn.bot/pkg/config"
	handler "iam43x.vpn.bot/pkg/handler"
)

func main() {

	bot, err := tgbotapi.NewBotAPI(config.Config.ApiToken)

	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for updateEvent := range updates {
		log.Printf("[%s] %s", updateEvent.Message.From.UserName, updateEvent.Message.Text)
		if updateEvent.Message.IsCommand() {
			handler.Handle(bot, updateEvent)
		} else {
			msg := tgbotapi.NewMessage(updateEvent.Message.Chat.ID, "Please use only commands!\n Get list commands /help")
			msg.ReplyToMessageID = updateEvent.Message.MessageID
			bot.Send(msg)
		}
	}
}
