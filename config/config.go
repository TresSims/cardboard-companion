package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var current *Config

var CfgFile string

func Load() *Config {
	if current != nil {
		return current
	}

	current = &Config{}
	viper.Unmarshal(current)
	return current
}

func InitConfig() {
	if CfgFile != "" {
		viper.SetConfigFile(CfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigType("yml")
		viper.SetConfigName("cardboardcompanion")
	}

	for key, value := range defaults {
		viper.SetDefault(key, value)
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Couldn't use config file:", viper.ConfigFileUsed())
	}
}
