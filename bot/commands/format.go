package commands

import (
	"cardboard-companion/bot/interactions"

	"github.com/bwmarrin/discordgo"
)

var formatCmd = &Definition{
	"format",
	&cmd{
		Description: "Ask folks about the format",
	},
	func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		response := interactions.PollInteraction.InteractionToInteractionResponse()
		s.InteractionRespond(i.Interaction, &response)
	},
}

func init() {
	Register(formatCmd)
}
