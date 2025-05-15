package bot

import (
	"cardboard-companion/bot/commands"

	"github.com/bwmarrin/discordgo"
)

type CommandHandler = func(s *discordgo.Session, i *discordgo.InteractionCreate)

func (b *bot) slashCommandRouter(s *discordgo.Session, i *discordgo.InteractionCreate) {
	name := i.ApplicationCommandData().Name
	_, handlers := commands.All()

	if handler, ok := handlers[name]; ok {
		handler(s, i)
	}
}

var registeredCommands []*discordgo.ApplicationCommand
