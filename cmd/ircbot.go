package main

import (
	"flag"
	"fmt"
	"strings"

	ircbot "github.com/recanman/irc-logbot/internal"
	"github.com/recanman/irc-logbot/packages/client"
)

func main() {
	// Define flags without default values
	serverPtr := flag.String("server", "", "IRC server address")
	portPtr := flag.Int("port", 0, "Port number")
	nicknamePtr := flag.String("nickname", "", "Nickname for the bot")
	channelPtr := flag.String("channels", "", "Comma-separated list of channels to join")
	fileNamePtr := flag.String("file", "log", "File name prefix for logging")

	// Parse flags
	flag.Parse()

	// Check if required flags are set
	if *serverPtr == "" || *portPtr == 0 || *nicknamePtr == "" || *channelPtr == "" || *fileNamePtr == "" {
		fmt.Println("Missing required flags. Please specify --server, --port, --nickname, --channels, and --file.")
		return
	}

	// Retrieve flag values
	server := *serverPtr
	port := *portPtr
	nickname := *nicknamePtr
	channelsStr := *channelPtr                  // Channels string separated by comma
	channels := strings.Split(channelsStr, ",") // Splitting the string into a slice

	// Create client with specified options
	client := client.Create(server, port, nickname, client.ClientOptions{
		Channels: channels,
	})

	fmt.Println("Connecting to server...")
	err := client.Connect()
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}

	ircbot.FromClient(client, *fileNamePtr)
	fmt.Println("Bot is running...")

	select {}
}
