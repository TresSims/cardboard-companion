package commands

import (
	"fmt"
	"math/rand"

	"github.com/bwmarrin/discordgo"
)

var wheelCmd = &Definition{
	"wheel",
	&cmd{
		Description: "Spin da wheel",
	},
	func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		g := rand.Intn(len(games))
		content := fmt.Sprintf(`It's time to spine the Wheel!
			You got: %s`, games[g])
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: content,
			},
		})
	},
}

var games = []string{
	"Planechase",
	"Share the Spoils",
	"Quiz Commander",
	"Truly Random",
	"Monochrome Matchup",
	"Mechanic Mashup",
	"Any # of Fools",
	"Bidget Battle",
	"Keywork Klash",
	"Partner Up",
	"Guild Wars",
	"Oathbreaker",
	"Tribal Throwdown",
	"Color Swap",
	"Pauper EDH",
	"Oldies but Goodies",
	"Framed Fight",
	"Alt Win Con",
	"Set Showdown",
	"Teeny Weenies",
	"Big Chungies",
	"Planeswalekr Party",
	"Precon Party",
	"Usurper/Kingdoms",
	"Archenemy",
	"Deck Swap",
	"Gifts Only",
	"Booster Pack Madness",
	"Bounties",
	"Two-Headed Giant",
}

func init() {
	Register(wheelCmd)
}
