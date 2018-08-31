package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	Token string
)

func init() {
	Token = os.Getenv("DISCORD_TOKEN")
}

func main() {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

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

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	langs := map[string]bool{"go": true}

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.Contains(m.Content, "+compilebot") {
		a := strings.Split(m.Content, " ")
		lang := a[1]

		if _, ok := langs[lang]; !ok {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s not supported.", lang))
			return
		}

		s.ChannelMessageSend(m.ChannelID, "Working..")

		code := strings.Split(m.Content, a[0]+" "+a[1])

		fmt.Println(code)

		f := strings.Replace(code[1], "`", "", -1)
		res, err := runnerClient(f)

		fmt.Println(f)

		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Error: "+err.Error())
			fmt.Println(f)
			return
		}

		s.ChannelMessageSend(m.ChannelID, res)

	}

}
