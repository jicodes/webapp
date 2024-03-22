package initializers

import (
	"os"

	"github.com/jicodes/webapp/models"
	"github.com/rs/zerolog"
)

func SyncDB () {

	logFile, logErr := os.OpenFile("/tmp/webapp.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if logErr != nil {
		panic(logErr)
	}
	defer logFile.Close()
	logger := zerolog.New(logFile).Level(zerolog.InfoLevel).With().Timestamp().Logger()

	if err := DB.AutoMigrate(&models.User{}); err != nil {
			logger.Error().Err(err).Msg("Error migrating the database schema")
	} else {
		logger.Info().Msg("Database schema migrated")
	}
}