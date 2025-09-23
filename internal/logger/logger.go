package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func NewLogger(log_type string) {
	zerolog.TimeFieldFormat = time.RFC3339
	switch log_type {
	case "prod":
		log.Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
	case "dev":
		log.Logger = zerolog.New(
			zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "15:04:05"},
		).With().Timestamp().Logger()
	default:
		log.Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
	}
}
