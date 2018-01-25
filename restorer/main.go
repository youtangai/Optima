package main

import (
	"flag"

	"github.com/gin-gonic/gin"
	"github.com/youtangai/Optima/restorer/config"
	"github.com/youtangai/Optima/restorer/controller"
)

const (
	//PORT is ポート番号
	PORT = "62073"
)

func main() {
	//コマンドライン引数を取得
	secretKeyPath := flag.String("i", "/var/optima/optima_key", "indentity file path for scp for controller node")
	controllerIP := flag.String("ip", "192.168.64.12", "controller node's ip")
	flag.Parse()

	//秘密鍵のパスを設定
	config.SetSecretKeyPath(*secretKeyPath)
	config.SetControllerIP(*controllerIP)

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.POST("/restore", controller.RestoreContainerController)
	router.Run(":" + PORT)
}
