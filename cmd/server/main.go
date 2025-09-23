package main

import (
	"github.com/gonzalocarreteroh/weather-api/config"
	"github.com/gonzalocarreteroh/weather-api/internal/api"
	"github.com/gonzalocarreteroh/weather-api/internal/cache"
	"github.com/gonzalocarreteroh/weather-api/internal/logger"
	"github.com/gonzalocarreteroh/weather-api/internal/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
)

func main() {
	cfg, _ := config.LoadConfig()
	logger.NewLogger(cfg.LogType)
	rdb := cache.NewRedisCache(cfg.RedisURL)
	svc := &service.WeatherService{Cache: rdb, APIKey: cfg.WeatherAPIKey}
	h := &api.Handler{Weather: svc}

	e := echo.New()

	// Middlewares
	e.Use(middleware.Recover())

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:  true,
		LogURI:     true,
		LogMethod:  true,
		LogLatency: true, // nice to see request timing
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			log.Info().
				Str("method", v.Method).
				Str("uri", v.URI).
				Int("status", v.Status).
				Dur("latency", v.Latency).
				Msg("request handled")
			return nil
		},
	}))

	e.GET("/weather/:city", h.GetWeather)

	log.Info().Msgf("Starting server on :%s", cfg.ServerPort)
	e.Logger.Fatal(e.Start(":" + cfg.ServerPort))
}
