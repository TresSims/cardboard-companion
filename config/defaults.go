package config

type Config struct {
	Token   string
	GuildID string

	PollSchedule string
	PollChannel  string
}

var defaults = map[string]any{}
