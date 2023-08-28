package logger

import (
	"os"

	config "github.com/gookit/config/v2"
	"github.com/rs/zerolog"
)

var AppLog zerolog.Logger

func InitLogger() {
	AppLog = zerolog.New(os.Stderr).
		Level(zerolog.Level(config.Int("logLevel"))).
		With().
		Timestamp().
		Caller().
		Logger()
}
