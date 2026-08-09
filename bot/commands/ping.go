package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
)

var pingCmd = &Definition{
	"ping",
	&cmd{
		Description: "A basic healthcheck",
	},
	func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "pong!",
			},
		})
		if err != nil {
			log.Error().Err(err).Msg("Couldn't call the format command")
		}
	},
}

func init() {
	Register(pingCmd)
}
