package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jicodes/webapp/initializers"
	"github.com/jicodes/webapp/internals/logger"
)

func CheckHealthz(c *gin.Context) {
	c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Header("Pragma", "no-cache")
	c.Header("X-Content-Type-Options", "nosniff")

	if err := initializers.DB.Exec("SELECT 1").Error; err == nil {
		c.Status(http.StatusOK)
		logger.Logger.Info().Msg("Health check passed: 200 OK")
	} else {
		c.Status(http.StatusServiceUnavailable)
		logger.Logger.Error().Err(err).Msg("Health check failed: 503 Service Unavailable")
	}
}