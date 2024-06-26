package main

import (
	"fmt"

	ircbot "github.com/recanman/irc-logbot/internal"
	"github.com/recanman/irc-logbot/packages/client"
)

func main() {
	server := "irc.libera.chat"
	port := 6667
	nickname := "LogBot"
	channel := "#monero-community"

	fileName := "log"

	client := client.Create(server, port, nickname, client.ClientOptions{
		Channels: []string{channel},
	})

	fmt.Println("Connecting to server...")
	err := client.Connect()
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}

	ircbot.FromClient(client, fileName)
	fmt.Println("Bot is running...")

	select {}
}
