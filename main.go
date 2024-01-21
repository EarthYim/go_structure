package main

import (
	"go_stucture/config"
	"go_stucture/core/health"
	"go_stucture/gin"
)

func main() {
	//Get Config
	config := config.GetConfig()

	//init handler
	heatlHandler := health.NewHealthHandler()

	//Init http server with GIN
	r := gin.NewGinRouter()

	//Register /health
	r.GET("/health", heatlHandler.Handle)

	r.Run(config.Http.Port)
}
