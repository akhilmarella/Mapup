package main

import (
	"mapup/handlers"
	"mapup/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	route.GET("/status", handlers.HealthCheck)

	route.GET("/token", handlers.GetToken)
	route.POST("/intersection", middleware.IsAuthorized(), handlers.GetIntersection)

	route.Run("localhost:8080")
}
