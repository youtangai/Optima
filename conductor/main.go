package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/youtangai/Optima/conductor/controller"
)

const (
	GIN_PORT = "62070"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.POST("/load_indicator", controller.RegistLoadIndicator)
	router.Run(":" + GIN_PORT)
}
