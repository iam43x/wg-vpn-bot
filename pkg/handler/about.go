package handler

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func About(bot *tgbotapi.BotAPI, updateEvent tgbotapi.Update) {
	msg := tgbotapi.NewMessage(updateEvent.Message.Chat.ID, "WireGuard vpn server bot.\nhttps://github.com/iam43x/wg-vpn-bot\n@iam43x")
	bot.Send(msg)
}
