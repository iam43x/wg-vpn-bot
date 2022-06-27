package command

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func Help(bot *tgbotapi.BotAPI, updateEvent tgbotapi.Update) {
	msg := tgbotapi.NewMessage(updateEvent.Message.Chat.ID, "WireGuard vpn server bot\n\n@iam43x")
	bot.Send(msg)
}