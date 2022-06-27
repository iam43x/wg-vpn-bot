package command

import (
	"fmt"
	"log"
	"iam43x.vpn.bot/pkg/client"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Execute(command string, bot *tgbotapi.BotAPI, updateEvent tgbotapi.Update) {
	clientName := updateEvent.Message.From.UserName
	switch command[1:] {
		case "get":
			client.Get(bot, updateEvent)
		case "add":
			client.Create(clientName)
			info := fmt.Sprintf("Client ticket for user [%v] was created! Contact @iam43x to get access :)", clientName)
			message := tgbotapi.NewMessage(updateEvent.Message.Chat.ID, info)
			bot.Send(message)
			log.Println(info)
		case "help":
			Help(bot, updateEvent)
		default:
			Help(bot, updateEvent)
	}
}