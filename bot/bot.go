package bot

import (
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

func (b *bot) init() {
	// TODO: Implement init
	b.initOnce.Do(func() {
		b.dg.Identify.Intents = discordgo.IntentsGuildMessages
		b.dg.Open()

		b.sc = make(chan os.Signal, 1)
		b.scheduler = cron.New()

		b.dg.AddHandler(b.handleMessages)
		// b.dg.AddHandler(b.slashCommandRouter)
	})
}

func (b *bot) destroy() {
	// TODO: Implement destroy
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
