package interactions

import "github.com/bwmarrin/discordgo"

var answers = []discordgo.PollAnswer{
	{
		AnswerID: 0,
		Media:    &discordgo.PollMedia{Text: "60 Card Format"},
	},
	{
		AnswerID: 1,
		Media:    &discordgo.PollMedia{Text: "EDH"},
	},
	{
		AnswerID: 2,
		Media:    &discordgo.PollMedia{Text: "The Wheel"},
	},
	{
		AnswerID: 3,
		Media:    &discordgo.PollMedia{Text: "Other/Don't Care"},
	},
	{
		AnswerID: 4,
		Media:    &discordgo.PollMedia{Text: "Can't Play Today"},
	},
}

var formatPoll = &discordgo.Poll{
	Question:         discordgo.PollMedia{Text: "Another Thursday? Another Poll!"},
	Answers:          answers,
	AllowMultiselect: false,
	Duration:         8,
}

var pollInteraction = &discordgo.InteractionResponse{
	Type: discordgo.InteractionResponseChannelMessageWithSource,
	Data: &discordgo.InteractionResponseData{
		Poll: formatPoll,
	},
}

func PollInteraction() *discordgo.InteractionResponse {
	return pollInteraction
}

var pollMessage = &discordgo.MessageSend{
	Poll: formatPoll,
}

func PollMessage() *discordgo.MessageSend {
	return pollMessage
}
