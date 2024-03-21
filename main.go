package main

import (
	"github.com/gin-gonic/gin"

	"github.com/jicodes/webapp/handlers"
	"github.com/jicodes/webapp/initializers"
	"github.com/jicodes/webapp/internals/logger"
	"github.com/jicodes/webapp/middlewares"
)

func init () {
	logger.InitLogger()
	
	initializers.LoadVariables()
	initializers.ConnectDB()
	initializers.SyncDB()
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.CheckRequestMethod())
	r.Use(middlewares.CheckPayload())

	r.GET("/healthz", handlers.CheckHealthz)
	r.POST("/v1/user", handlers.CreateUser)
	r.GET("/v1/user/self", middlewares.BasicAuth(), handlers.GetUser)
	r.PUT("/v1/user/self", middlewares.BasicAuth(), handlers.UpdateUser)

	return r
}

func main() {
	r := setupRouter()
	r.Run()
}