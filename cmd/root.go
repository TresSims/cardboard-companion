package cmd

import (
	"cardboard-companion/config"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cardboard",
	Short: "Cardboard Companion is a go bot and Next.JS web UI",
	Long: `A discord bot written in go, with a web ui written in Next.
					Tailored to managing board game servers (MTG In particular)`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal().Err(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&config.CfgFile, "config", "", "config file (default is ./cardboard-companion.yml)")

	cobra.OnInitialize(config.InitConfig)
}
