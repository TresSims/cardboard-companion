package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
)

func (b *bot) handleMessages(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	msg := fmt.Sprintf("RCV [%v] - %v", m.ChannelID, m.Content)
	log.Info().Msg(msg)
}
