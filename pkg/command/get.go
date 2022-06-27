package command

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"iam43x.vpn.bot/pkg/config"
	"iam43x.vpn.bot/pkg/security"
)

func Get(bot *tgbotapi.BotAPI, updateEvent tgbotapi.Update) {
	clientName := updateEvent.Message.From.UserName
	if security.AllowUser(clientName) {
		message := tgbotapi.NewMessage(updateEvent.Message.Chat.ID, "*** Access denied! ***")
		bot.Send(message)
	} else {
		configPath := config.Config.ClientConfig[clientName]
		configFile := tgbotapi.NewInputMediaDocument(tgbotapi.FilePath(configPath))
		mediaGroup := tgbotapi.NewMediaGroup(updateEvent.Message.Chat.ID, []interface{}{
			configFile,
		})
		bot.SendMediaGroup(mediaGroup)
	}

}
