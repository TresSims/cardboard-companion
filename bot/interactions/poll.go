package interactions

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

var pollDuration = 10

var minutesUntilGame int = pollDuration*60 + 30

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
		Media:    &discordgo.PollMedia{Text: "Jack Goes to Selesnya Jail"},
	},
	{
		AnswerID: 5,
		Media:    &discordgo.PollMedia{Text: "Can't Play Today"},
	},
}

var formatPoll = discordgo.Poll{
	Question:         discordgo.PollMedia{Text: "Another Thursday? Another Poll!"},
	Answers:          answers,
	AllowMultiselect: false,
	Duration:         pollDuration,
}

func buildPollText(minutesLeft int) string {
	now := time.Now()

	gameTime := int(now.Unix()) + minutesLeft*60 // minutes -> seconds

	return fmt.Sprintf("It's game time folks! Please vote in the poll, it's how we know who's coming to play! See you at <t:%d>", gameTime)
}

var PollInteraction = &Interaction{
	Content:         func() string { return buildPollText(minutesUntilGame) },
	Poll:            func() *discordgo.Poll { return &formatPoll },
	Embeds:          func() []*discordgo.MessageEmbed { return nil },
	TTS:             func() bool { return false },
	Components:      func() []discordgo.MessageComponent { return nil },
	Files:           func() []*discordgo.File { return nil },
	AllowedMentions: func() *discordgo.MessageAllowedMentions { return nil },
}
