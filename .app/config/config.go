package config

import (
	"github.com/spf13/viper"
)

func Setup(envPath string) {
	viper.SetConfigFile(envPath)
	viper.AutomaticEnv()
	viper.ReadInConfig()

}
