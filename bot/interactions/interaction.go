package interactions

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
)

type Interaction struct {
	Content         func() string
	Embeds          func() []*discordgo.MessageEmbed
	TTS             func() bool
	Components      func() []discordgo.MessageComponent
	Files           func() []*discordgo.File
	AllowedMentions func() *discordgo.MessageAllowedMentions
	Poll            func() *discordgo.Poll
}

func (i Interaction) InteractionToInteractionResponse() discordgo.InteractionResponse {
	log.Debug().Msg(fmt.Sprintf("Building interaction response for %T", i))
	return discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content:         i.Content(),
			Embeds:          i.Embeds(),
			Poll:            i.Poll(),
			TTS:             i.TTS(),
			Components:      i.Components(),
			Files:           i.Files(),
			AllowedMentions: i.AllowedMentions(),
		},
	}
}

func (i Interaction) InteractionToMessageSend() discordgo.MessageSend {
	log.Debug().Msg(fmt.Sprintf("Building message send for %T", i))
	return discordgo.MessageSend{
		Content:         i.Content(),
		Poll:            i.Poll(),
		Embeds:          i.Embeds(),
		TTS:             i.TTS(),
		Components:      i.Components(),
		Files:           i.Files(),
		AllowedMentions: i.AllowedMentions(),
	}
}
