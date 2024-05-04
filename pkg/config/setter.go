package config

import (
	"log"
	"github.com/spf13/viper"
)

func Set() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Config not found")
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatal("config error")
	}
}
