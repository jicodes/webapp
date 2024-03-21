package logger

import (
	"os"

	"github.com/rs/zerolog"
)

var logger zerolog.Logger

func InitLogger() {
	file, err := os.OpenFile("/tmp/webapp.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	logger = zerolog.New(file).With().Timestamp().Logger()
	logger.Info().Msg("Logger initialized")
}

func GetLogger() zerolog.Logger {
	return logger
}