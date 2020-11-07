package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var BotID string

func main() {
	// New bot instance
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Grab bot user info
	u, err := dg.User("@me")
	if err != nil {
		fmt.Println(err.Error())
	}

	// BotID from user info
	BotID = u.ID

	// Chatlog listener
	dg.AddHandler(messageHandler)

	// Open connection with discord
	err = dg.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Print Bot startup status and open a channel to keep bot running
	fmt.Println("Bot " + BotID + "is running")
	<-make(chan struct{})
	return
}

// messageHandler listens to chat
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// If the message author is the bot, do nothing
	if m.Author.ID == BotID {
		return
	}

	if m.Content == "ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
	}

}
