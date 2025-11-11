package config

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Config struct {
	ServerPort    string `mapstructure:"SERVER_PORT"`
	WeatherAPIKey string `mapstructure:"WEATHER_API_KEY"`
	RedisURL      string `mapstructure:"REDIS_URL"`
	LogType       string `mapstructure:"LOG_TYPE"`
}

func LoadConfig() (Config, error) {
	var cfg Config
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Warn().Msg("Error reading config file in config.go")
		// return cfg, err
	}
	err := viper.Unmarshal(&cfg)
	log.Info().Msg("Config loaded successfully")
	log.Info().Msgf("Config RedisURL: %+v", cfg.RedisURL)
	return cfg, err
}
