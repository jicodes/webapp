package logger

import (
	"os"

	"github.com/rs/zerolog"
)

var Logger zerolog.Logger

func InitLogger() {
	file, err := os.OpenFile("/tmp/webapp.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	Logger = zerolog.New(file).With().Timestamp().Logger()
	Logger.Info().Msg("Logger initialized")
}