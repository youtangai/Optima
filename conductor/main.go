package main

import (
	"github.com/gin-gonic/gin"
	"github.com/youtangai/Optima/conductor/config"
	"github.com/youtangai/Optima/conductor/controller"
)

func main() {
	port := config.GinPort()
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.POST("/load_indicator", controller.RegistLoadIndicator)
	router.Run(":" + port)
}
