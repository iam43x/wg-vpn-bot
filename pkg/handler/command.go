package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"iam43x.vpn.bot/pkg/command"
)

func Handle(bot *tgbotapi.BotAPI, updateEvent tgbotapi.Update) {
	switch updateEvent.Message.Command() {
	case "get":
		command.Get(bot, updateEvent)
	case "add":
		command.Add(bot, updateEvent)
	case "about":
		About(bot, updateEvent)
	default:
		msg := tgbotapi.NewMessage(updateEvent.Message.Chat.ID, "Pls check cmd list.")
		msg.ReplyToMessageID = updateEvent.Message.MessageID
		bot.Send(msg)
	}
}
