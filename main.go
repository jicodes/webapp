package main

import (
	"github.com/gin-gonic/gin"

	"github.com/jicodes/webapp/handlers"
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

	r.Any("/healthz", middlewares.CheckRequestMethod(), middlewares.CheckPayload(), handlers.CheckHealthz)
	r.POST("/v2/user", handlers.CreateUser)
	r.GET("/v1/user/verify/", handlers.VerifyEmail)

	userGroup := r.Group("/v2/user")
	userGroup.Use(middlewares.BasicAuth())  
	userGroup.Use(middlewares.NeedVerify())  
	{
		userGroup.GET("/self", handlers.GetUser)
		userGroup.PUT("/self", handlers.UpdateUser)
	}

	return r
}

func main() {
	r := setupRouter()
	r.Run()
}