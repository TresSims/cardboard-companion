package config

type Config struct {
	Token   string
	GuildID string

	PollSchedule string
	PollChannel  string

	Web struct {
		Port     int
		Frontend string
	}
}

var defaults = map[string]any{}
