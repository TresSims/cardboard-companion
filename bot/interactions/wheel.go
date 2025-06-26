package interactions

import (
	"fmt"
	"math/rand"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
)

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

func buildWheelContent() string {
	g := rand.Intn(len(games))
	log.Info().Msg("Sending Wheel Interaction")
	return fmt.Sprintf(`It's time to spine the Wheel!
			You got: %s`, games[g])
}

var wheelInteraction = &discordgo.InteractionResponse{
	Type: discordgo.InteractionResponseChannelMessageWithSource,
	Data: &discordgo.InteractionResponseData{
		Content: buildWheelContent(),
	},
}

func WheelInteraction() *discordgo.InteractionResponse {
	return wheelInteraction
}
