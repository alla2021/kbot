package main

import (
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	tele "gopkg.in/telebot.v3"
)

func main() {
	// Initializing root command using Cobra
	var rootCmd = &cobra.Command{
		Use:   "kbot",
		Short: "Telegram bot with root command",
		Run: func(cmd *cobra.Command, args []string) {
			// Starting the Telegram bot
			startTelegramBot()
		},
	}

	// Adding root command to Cobra
	var cmdKbot = &cobra.Command{Use: "kbot"}
	rootCmd.AddCommand(cmdKbot)

	// Calling the root command
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func startTelegramBot() {
	// Getting bot token from environment variable
	teleToken :=  os.Getenv("TELE_TOKEN")
	if teleToken == "" {
		log.Fatal("Bot token value not set for Telegram.")
	}

	// Telegram bot settings
	pref := tele.Settings{
		Token:  teleToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	// Creating a new Telegram bot object
	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	// Handler for the /start command
	b.Handle("/start", func(c tele.Context) error {
		return c.Send("Welcome! This is the start command.")
	})

	// Handler for the /help command
	b.Handle("/help", func(c tele.Context) error {
		return c.Send("This is the help command. Here you can get help.")
	})

	// Handler for the /info command
	b.Handle("/info", func(c tele.Context) error {
		return c.Send("This is the info command. Here you can get information.")
	})

	// Starting the Telegram bot
	b.Start()
}



