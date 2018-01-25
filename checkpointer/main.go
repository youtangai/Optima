package main

import (
	"github.com/gin-gonic/gin"
	"github.com/youtangai/Optima/checkpointer/controller"
)

const (
	//PORT is ポート番号
	PORT = "62072"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.POST("/checkpoint", controller.CehckpointContainer)
	router.Run(":" + PORT)
}
