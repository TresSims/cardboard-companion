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
		s.InteractionRespond(i.Interaction, interactions.PollInteraction())
	},
}

func init() {
	Register(formatCmd)
}
