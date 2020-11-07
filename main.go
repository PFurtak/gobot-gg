package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

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
	BotID := u.ID

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
