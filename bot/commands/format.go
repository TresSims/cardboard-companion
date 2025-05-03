package commands

import "github.com/bwmarrin/discordgo"

var answers = []discordgo.PollAnswer{
	{
		AnswerID: 0,
		Media:    &discordgo.PollMedia{Text: "EDH"},
	},
}

var FormatPoll = &discordgo.Poll{
	Question:         discordgo.PollMedia{Text: "Format?"},
	Answers:          answers,
	AllowMultiselect: false,
	Duration:         100,
}

var formatCmd = &Definition{
	"format",
	&cmd{
		Description: "Ask folks about the format",
	},
	func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Poll: FormatPoll,
			},
		})
	},
}

func init() {
	Register(formatCmd)
}
