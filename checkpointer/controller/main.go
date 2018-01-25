package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CehckpointContainer(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "checkpoint",
	})
}
