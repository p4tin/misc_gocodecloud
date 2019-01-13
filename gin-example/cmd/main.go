package main

import (
	"fmt"
	example "misc/gin-example"
	"misc/gin-example/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	healthHandler := handlers.CreateService(example.Config{
		BuildNumber: 1,
		Environment: "Test",
	})

	router := gin.Default()

	router.GET("/health", healthHandler.HealthHandler)

	err := router.Run(":12345")
	if err != nil {
		panic(fmt.Sprintf("Webserver stopped with error %s, quitting application", err.Error()))
	}

}
