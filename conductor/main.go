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
	router.POST("/create_dir", controller.CreateDirController)
	router.POST("/join", controller.JoinController)
	router.POST("/leave", controller.LeaveController)
	router.Run(":" + GIN_PORT)
}
