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
	r.Use(middlewares.CheckRequestMethod())
	r.Use(middlewares.CheckPayload())

	r.GET("/healthz", handlers.CheckHealthz)
	r.POST("/v1/user", handlers.CreateUser)

	authGroup := r.Group("/v1/user/self")
	authGroup.Use(middlewares.BasicAuth())  
	authGroup.Use(middlewares.NeedVerify()) 
	{
		authGroup.GET("", handlers.GetUser)
		authGroup.PUT("", handlers.UpdateUser)
	}

	return r
}

func main() {
	r := setupRouter()
	r.Run()
}