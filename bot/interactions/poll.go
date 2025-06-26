package interactions

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

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
		Media:    &discordgo.PollMedia{Text: "Other/Don't Care"},
	},
	{
		AnswerID: 3,
		Media:    &discordgo.PollMedia{Text: "Can't Play Today"},
	},
}

var formatPoll = &discordgo.Poll{
	Question:         discordgo.PollMedia{Text: "Another Thursday? Another Poll!"},
	Answers:          answers,
	AllowMultiselect: false,
	Duration:         10,
}

var pollInteraction = &discordgo.InteractionResponse{
	Type: discordgo.InteractionResponseChannelMessageWithSource,
	Data: &discordgo.InteractionResponseData{
		Poll:    formatPoll,
		Content: buildPollText(formatPoll.Duration*60 + 30), // hours -> minutes
	},
}

func buildPollText(minutesLeft int) string {
	now := time.Now()

	gameTime := int(now.Unix()) + minutesLeft*60 // minutes -> seconds

	return fmt.Sprintf("It's game time folks! Please vote in the poll, it's how we know who's coming to play! See you at <t:%d>", gameTime)
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
