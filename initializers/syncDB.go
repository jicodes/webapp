package initializers

import (
	"github.com/jicodes/webapp/models"
)

func SyncDB () {
	if err := DB.AutoMigrate(&models.User{}); err != nil {
			initializersLogger.Error().Err(err).Msg("Error migrating the database schema")
	} else {
		initializersLogger.Info().Msg("Database schema migrated")
	}
}