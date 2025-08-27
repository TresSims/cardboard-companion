package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
)

func (b *bot) send(channel string, message *discordgo.MessageSend) {
	_, err := b.dg.ChannelMessageSendComplex(channel, message)
	if err != nil {
		log.Error().Err(err).Msg("Couldn't run scheduled task!")
	}
}
