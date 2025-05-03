package commands

import "github.com/bwmarrin/discordgo"

var wheelCmd = &Definition{
	"wheel",
	&cmd{
		Description: "Spin da wheel",
	},
	func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: `It's time to spin the wheel!
				You got, Booster Pack Maddness!`,
			},
		})
	},
}

func init() {
	Register(wheelCmd)
}
