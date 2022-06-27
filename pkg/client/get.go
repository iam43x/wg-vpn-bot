package client

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"iam43x.vpn.bot/pkg/config"
)

func Get(bot *tgbotapi.BotAPI, updateEvent tgbotapi.Update) {
	clientName := updateEvent.Message.From.UserName
	configPath := config.Config.ClientConfig[clientName]
	configFile := tgbotapi.NewInputMediaDocument(tgbotapi.FilePath(configPath))
	mediaGroup := tgbotapi.NewMediaGroup(updateEvent.Message.Chat.ID, []interface{}{
		configFile,
	})
	bot.SendMediaGroup(mediaGroup)
}