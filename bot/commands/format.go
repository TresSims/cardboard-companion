package commands

import (
	"cardboard-companion/bot/interactions"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
)

var formatCmd = &Definition{
	"format",
	&cmd{
		Description: "Ask folks about the format",
	},
	func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		response := interactions.PollInteraction.InteractionToInteractionResponse()
		err := s.InteractionRespond(i.Interaction, &response)
		if err != nil {
			log.Error().Err(err).Msg("Couldn't call the format command")
		}
	},
}

func init() {
	Register(formatCmd)
}
