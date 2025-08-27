package config

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

var current *Config

var CfgFile string

func Load() *Config {
	if current != nil {
		return current
	}

	current = &Config{}
	err := viper.Unmarshal(current)
	if err != nil {
		log.Error().Err(err).Msg("Couldn't unmarchal config")
	}
	return current
}

func InitConfig() {
	if CfgFile != "" {
		viper.SetConfigFile(CfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigType("yml")
		viper.SetConfigName("cardboard-companion")
	}

	for key, value := range defaults {
		viper.SetDefault(key, value)
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Error().Err(err).Msg("Could not read in config.")
	}
}
