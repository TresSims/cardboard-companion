package models

import (
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model

	Schedule    string
	Type        string
	Interaction string
}

func (t *Task) Run() {
	switch t.Type {
	case "sendInteraction":
		// TODO:
		// sendInteraction(t.Interaction)
		return
	default:
		log.Error().Msg("No task of specified type")
		return
	}
}

func init() {
	AllModels = append(AllModels, Task{})
}
