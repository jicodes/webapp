package initializers

import (
	"github.com/jicodes/webapp/models"
	"github.com/jicodes/webapp/internals/logger"
)

func SyncDB () {
	if err := DB.AutoMigrate(&models.User{}); err != nil {
			logger.Logger.Error().Err(err).Msg("Error migrating the database schema")
	} else {
		logger.Logger.Info().Msg("Database schema migrated")
	}
}