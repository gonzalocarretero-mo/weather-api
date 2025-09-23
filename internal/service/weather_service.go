package service

import (
	"context"
	"encoding/json"
	"time"

	"github.com/gonzalocarreteroh/weather-api/internal/cache"
	"github.com/gonzalocarreteroh/weather-api/internal/client"
	"github.com/rs/zerolog/log"
)

type WeatherService struct {
	Cache  *cache.RedisCache
	APIKey string
}

func (s *WeatherService) GetWeather(ctx context.Context, city string) (client.WeatherResponse, error) {
	// Check if data already in cache
	if val, err := s.Cache.Get(ctx, city); err == nil {
		var wr client.WeatherResponse
		if err := json.Unmarshal([]byte(val), &wr); err != nil {
			log.Warn().
				Msg("Couldn't unpack cache into WeatherResponse type")
		}
		log.Info().Msgf("Data for %s was in cache. Returning cached data.", city)
		return wr, nil
	}

	// Fetch from API
	log.Info().Msgf("Data for %s was not in cache. Fetching from API.", city)
	wr, err := client.FetchWeather(s.APIKey, city)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to fetch weather data for %s", city)
		return wr, err
	}

	// Save in cache
	bytes, _ := json.Marshal(wr)
	if err := s.Cache.Set(ctx, city, string(bytes), 30*time.Minute); err != nil {
		log.Warn().Err(err).Msg("Couldn't save data in cache")
	}
	log.Info().Msgf("Data for %s cached successfully.", city)

	return wr, nil
}
