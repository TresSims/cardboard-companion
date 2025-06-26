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
		s.InteractionRespond(i.Interaction, interactions.WheelInteraction())
	},
}

func init() {
	Register(wheelCmd)
}
