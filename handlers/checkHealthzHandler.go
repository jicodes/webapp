package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jicodes/webapp/initializers"
)

func CheckHealthz(c *gin.Context) {
	c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Header("Pragma", "no-cache")
	c.Header("X-Content-Type-Options", "nosniff")

	if err := initializers.DB.Exec("SELECT 1").Error; err == nil {
		c.Status(http.StatusOK)
		log.Println("Health check passed: 200 OK, DB is up and running")
	} else {
		c.Status(http.StatusServiceUnavailable)
		log.Fatal("Health check failed: 503 Service Unavailable, DB is down or not reachable")
	}
}