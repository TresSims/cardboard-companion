package commands

import (
	"cardboard-companion/bot/interactions"

	"github.com/bwmarrin/discordgo"
)

var wheelCmd = &Definition{
	"wheel",
	&cmd{
		Description: "Spin da wheel",
	},
	func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		response := interactions.WheelInteraction.InteractionToInteractionResponse()
		s.InteractionRespond(i.Interaction, &response)
	},
}

func init() {
	Register(wheelCmd)
}
