package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jicodes/webapp/controllers"
	"github.com/jicodes/webapp/initializers"
	"github.com/jicodes/webapp/middlewares"
)

func init () {
	initializers.LoadVariables()
	initializers.ConnectDB()
	initializers.SyncDB()
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.CheckRequestMethod())
	r.Use(middlewares.CheckPayload())

	r.GET("/healthz", func(c *gin.Context) {
		c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
		c.Header("Pragma", "no-cache")
		c.Header("X-Content-Type-Options", "nosniff")

		if err := initializers.DB.Exec("SELECT 1").Error; err == nil {
			c.Status(http.StatusOK)
		} else {
			c.Status(http.StatusServiceUnavailable)
		}
	})
	
	r.POST("/v1/user", controllers.CreateUser)
	r.GET("/v1/user/self", middlewares.BasicAuth(), controllers.GetUser)
	r.PUT("/v1/user/self", middlewares.BasicAuth(), controllers.UpdateUser)

	return r
}

func main() {
	r := setupRouter()
	r.Run()
}