package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	Token string
	users *Users
)

func init() {
	flag.StringVar(&Token, "t", "", "Discord Bot Token")
	flag.Parse()
}

func main() {

	// Open data.json file
	dataJsonFile, err := os.Open("data.json")
	if err != nil {
		fmt.Println("Error while opening data.json file.")
	}

	// defers the closing of the file so we can access it later.
	defer dataJsonFile.Close()

	users, err = loadUsers(dataJsonFile)
	if err != nil {
		fmt.Println("Error while reading data.json file.")
	}

	// Create a new discord bot
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// Called everytime when a message is created in a channel the both has access
// to. Uses the AddHandler event.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// ignore t he bot's messages
	if m.Author.ID == s.State.User.ID {
		return
	}

	// If the message is !db respond with that users database details.
	if strings.EqualFold("!db", m.Content) {
		// get the user that created the message
		user := users.GetUser(m.Author.ID)

		// if the user is not found in data.json, tell them to contact administrator
		if user == nil {
			s.ChannelMessageSend(
				m.ChannelID,
				"User not found, please contact an administrator!",
			)
			return
		}

		// send the user a private message with their details
		dm, err := s.UserChannelCreate(m.Author.ID)
		if err != nil {
			fmt.Println("error creating channel:", err)
			s.ChannelMessageSend(
				m.ChannelID,
				"Something went wrong while sending the DM!",
			)
			return
		}

		// send database information through dms
		embed := GenerateDBEmbed(user)
		_, err = s.ChannelMessageSendEmbed(dm.ID, embed)
		if err != nil {
			fmt.Println("error sending DM message:", err)
			s.ChannelMessageSend(
				m.ChannelID,
				"Failed to send you a DM. "+
					"Did you disable DM in your privacy settings?",
			)
		}

		s.ChannelMessageSend(
			m.ChannelID,
			"Check your messages, I sent your database information!",
		)
		return

	}
}
