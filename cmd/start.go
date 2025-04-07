package cmd

import (
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
