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

	initOnce     sync.Once
	destroyOnece sync.Once
}

func (b *bot) init() {
	// TODO: Implement init
	return
}

func (b *bot) destroy() {
	// TODO: Implement destroy
	return
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
