package commands

import (
	"cardboard-companion/bot/interactions"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
)

var wheelCmd = &Definition{
	"wheel",
	&cmd{
		Description: "Spin da wheel",
	},
	func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		response := interactions.WheelInteraction.InteractionToInteractionResponse()
		err := s.InteractionRespond(i.Interaction, &response)
		if err != nil {
			log.Error().Err(err).Msg("Couldn't run the wheel command")
		}
	},
}

func init() {
	Register(wheelCmd)
}
