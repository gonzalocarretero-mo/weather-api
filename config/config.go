package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	ServerPort    string `mapstructure:"server_port"`
	WeatherAPIKey string `mapstructure:"WEATHER_API_KEY"`
	RedisURL      string `mapstructure:"REDIS_URL"`
	LogType       string `mapstructure:"LOG_TYPE"`
}

func LoadConfig() (Config, error) {
	var cfg Config
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		return cfg, err
	}
	err := viper.Unmarshal(&cfg)
	return cfg, err
}
