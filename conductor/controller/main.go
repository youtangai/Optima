package controller

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/youtangai/Optima/conductor/db"
	"github.com/youtangai/Optima/conductor/model"
	"github.com/youtangai/Optima/conductor/util"
)

const (
	PublicKeyName = "optima_key.pub"
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

func CreateDirController(c *gin.Context) {
	//ディレクトリを作る
	json := new(model.JoinJson)
	err := c.ShouldBindJSON(json)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	hostName := json.HostName
	err = os.Chdir("/var")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	err = os.MkdirAll("optima/"+hostName, 0777)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func JoinController(c *gin.Context) {

	json := new(model.JoinJson)
	err := c.ShouldBindJSON(json)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	hostName := json.HostName
	err = os.Chdir("/var/optima/" + hostName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	if bool := util.FileExists(PublicKeyName); bool {
		//公開鍵が存在したら
		initialJoin(hostName)
	}
	//公開鍵が存在しなかったら 再配置処理へ
}

func initialJoin(hostName string) error {
	// /root/.ssh/authorised_keyに追記
	// 公開鍵の削除
	// /etc/hostsにipとエイリアスを記述
	log.Println("pub key is exists")
	return nil
}
