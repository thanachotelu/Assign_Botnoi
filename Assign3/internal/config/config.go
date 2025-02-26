package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ChannelSecret      string
	ChannelAccessToken string
	Port               string
}

// LoadConfig อ่านค่าจาก .env และ return ค่า Config
func LoadConfig() *Config {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error loading config file: %v", err)
	}

	return &Config{
		ChannelSecret:      viper.GetString("LINE_CHANNEL_SECRET"),
		ChannelAccessToken: viper.GetString("LINE_CHANNEL_ACCESS_TOKEN"),
		Port:               viper.GetString("PORT"),
	}
}
