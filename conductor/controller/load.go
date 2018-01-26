package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/youtangai/Optima/conductor/db"
	"github.com/youtangai/Optima/conductor/model"
)

// RegistLoadIndicator is 負荷指標を登録するメソッド
func RegistLoadIndicator(c *gin.Context) {
	json := new(model.LoadIndicatorJson)
	err := c.ShouldBindJSON(&json)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	err = db.RegistLoadIndicator(*json)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
