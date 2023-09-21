package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	sess, err := discordgo.New("Bot MTE1MzczOTI4Nzk2NDE1NTk0NQ.G2mv89.ByL_7RFjdUJii-ZVhhhSMvdhUwT9V3iB44rKWA")
	if err != nil {
		log.Fatal(err)
	}
	sess.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {

		// Ignore the bot messages
		if m.Author.ID == s.State.User.ID {
			return
		}

		// if m.Author.ID == "270589534591385600" {
		// 	s.ChannelMessageSend(m.ChannelID, "Shut the fuck up Andre!")
		// }

		// Created a switch case that will look for specific messages from a specific user and give replies
		switch {
		case m.Author.ID == "606638110087839780":
			s.ChannelMessageSend(m.ChannelID, "Shut the fuck up Andre!")
		case m.Author.ID == "606638110087839780":
			s.ChannelMessageSend(m.ChannelID, "Oh Hey Mango :wave:")
		}

	})

	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged
	err = sess.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	fmt.Println("The Bot Is Running!")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

}
