package command

import (
	"fmt"
	"log"
	"os/exec"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type ClientConfigWgData struct {
	PublicKey  string
	AllowedIPs string
}

//TODO
func Add(bot *tgbotapi.BotAPI, updateEvent tgbotapi.Update) {
	clientName := updateEvent.Message.From.UserName
	cmd := exec.Command("sh", "-c", fmt.Sprintf("echo '%v' >> /etc/wireguard/users.txt", clientName))
	err := cmd.Run()
	if err != nil {
		log.Panicln(err)
	}
	info := fmt.Sprintf("Client ticket for user [%v] was created! Contact @iam43x to get access :)", clientName)
	message := tgbotapi.NewMessage(updateEvent.Message.Chat.ID, info)
	bot.Send(message)
	log.Println(info)
}
