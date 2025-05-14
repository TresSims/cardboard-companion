package bot

import (
	"github.com/bwmarrin/discordgo"
)

func (b *bot) send(channel string, message *discordgo.MessageSend) {
	b.dg.ChannelMessageSendComplex(channel, message)
}
