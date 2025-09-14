package cmd

import (
	"cardboard-companion/config"
	"cardboard-companion/server"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.AddCommand(serverStartCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Interact with the server",
	Long: `Commands for interacting with the Bot's server. 
					The bot they don't work well without the bot running.`,
}

var serverStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the bot's server",
	Run:   startserver,
}

func startserver(cmd *cobra.Command, args []string) {
	log.Info().Msg("Starting up the server")

	router := server.NewRouter()
	conf := config.Load()
	addr := fmt.Sprintf(":%d", conf.Web.Port)
	srv := &http.Server{Addr: addr, Handler: router}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("Server listen error.")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("Server shutdown failure")
	}

	select {
	case <-ctx.Done():
		log.Info().Msg("Server exiting!")
	case <-time.After(5 * time.Second):
		log.Warn().Msg("Server shutdown timeout hit, force-killing")
	}
}
