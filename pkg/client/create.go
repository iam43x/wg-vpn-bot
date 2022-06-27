package client

import (
	"fmt"
	"log"
	"os/exec"
)

type ClientConfigWgData struct {
	PublicKey string
	AllowedIPs string
}

func Create(name string) {
	//TODO
	cmd := exec.Command("sh", "-c", fmt.Sprintf("echo '%v' >> /etc/wireguard/users.txt", name))
	err := cmd.Run()
	if err != nil {
		log.Panicln(err)
	}
}