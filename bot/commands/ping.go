package commands

import "github.com/bwmarrin/discordgo"

var pingCmd = &Definition{
	"ping",
	&cmd{
		Description: "Test command to make sure the bots alive",
	},
	func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Pong!",
			},
		})
	},
}

func init() {
	Register(pingCmd)
}
