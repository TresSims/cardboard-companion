package cmd

import (
	"cardboardcompanion/bot"
	"cardboardcompanion/config"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(botCmd)
	botCmd.AddCommand(botStartCmd)
}

var botCmd = &cobra.Command{
	Use:   "bot",
	Short: "Interact with the bot",
	Long: `Commands for interacting with the bot, 
						e.g. Start, Stop`,
}

var botStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the bot",
	Long: `Command for starting the bot
					e.g. Start the bot`,
	Run: startBot,
}

func startBot(cmd *cobra.Command, args []string) {
	log.Info().Msg("Starting the bot!")

	conf := config.Load()
	token := conf.Token

	if token == "" {
		log.Fatal().Msg("No token available! Can't Start!")
	}

	b, err := bot.New(token)
	if err != nil {
		log.Fatal().Err(err).Msg("Error connecting to discord")
	}

	err = b.Start()
	if err != nil {
		log.Fatal().Err(err).Msg("Error starting Cardboard Companion!")
	}
	log.Info().Msg("Your friend has arrived!")

	b.Wait()
	b.Stop()
}
