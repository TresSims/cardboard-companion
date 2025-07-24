package bot

import (
	"cardboard-companion/bot/commands"
	"cardboard-companion/bot/interactions"
	"cardboard-companion/config"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/robfig/cron/v3"
	"github.com/rs/zerolog/log"
)

type Bot interface {
	Start() error

	Wait()

	Stop()
}

func New(token string) (Bot, error) {
	dg, err := discordgo.New("Bot " + token)
	b := &bot{dg: dg}
	b.init()

	return b, err
}

type bot struct {
	dg *discordgo.Session

	sc chan os.Signal

	scheduler *cron.Cron

	initOnce    sync.Once
	destroyOnce sync.Once
}

var conf *config.Config

func (b *bot) init() {
	b.initOnce.Do(func() {
		conf := config.Load()

		b.dg.Identify.Intents = discordgo.IntentsGuildMessages
		b.dg.Open()

		b.sc = make(chan os.Signal, 1)
		b.scheduler = cron.New()

		// Schedule Cron Jobs
		if conf.PollSchedule != "" {
			_, err := b.scheduler.AddFunc(conf.PollSchedule, func() {
				message := interactions.PollInteraction.InteractionToMessageSend()
				b.send(conf.PollChannel, &message)
			})
			if err != nil {
				log.Fatal().Err(err).Msg("Failed to start scheduler")
			}
			log.Info().Msg(fmt.Sprintf("Registering command at %s", conf.PollSchedule))
			b.scheduler.Start()
		}

		// Respond to text mesages
		b.dg.AddHandler(b.handleMessages)

		// Resgister Slash Commands
		b.dg.AddHandler(b.slashCommandRouter)
		cmdList, _ := commands.All()
		registeredCommands = make([]*discordgo.ApplicationCommand, len(cmdList))
		for idx, rawCmd := range cmdList {
			cmd, err := b.dg.ApplicationCommandCreate(b.dg.State.User.ID, conf.GuildID, rawCmd)
			if err != nil {
				log.Fatal().Err(err).Msg("Failed to register application commands")
			}
			registeredCommands[idx] = cmd
		}
	})
}

func (b *bot) destroy() {
	b.destroyOnce.Do(func() {
		b.scheduler.Stop()
	})
}

func (b *bot) Start() error {
	b.init()

	signal.Notify(b.sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGABRT, os.Interrupt)
	return nil
}

func (b *bot) Wait() {
	b.init()

	<-b.sc
}

func (b *bot) Stop() {
	b.init()

	log.Info().Msg("Draining down poliely. Please Wait")
	b.destroy()
	b.dg.Close()
	log.Info().Msg("All done!")
}
