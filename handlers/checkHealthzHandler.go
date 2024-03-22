package handlers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jicodes/webapp/initializers"
	"github.com/rs/zerolog"
)

func CheckHealthz(c *gin.Context) {
	logFile, err := os.OpenFile("/tmp/webapp.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()
	logger := zerolog.New(logFile).Level(zerolog.InfoLevel).With().Timestamp().Logger()

	c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Header("Pragma", "no-cache")
	c.Header("X-Content-Type-Options", "nosniff")

	if err := initializers.DB.Exec("SELECT 1").Error; err == nil {
		c.Status(http.StatusOK)
		logger.Info().Msg("Health check passed: 200 OK")
	} else {
		c.Status(http.StatusServiceUnavailable)
		logger.Error().Err(err).Msg("Health check failed: 503 Service Unavailable")
	}
}