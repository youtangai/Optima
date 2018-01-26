package controller

import (
	"net/http"
	"os"

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
	//公開鍵が存在したら initialJoin
	//公開鍵が存在しなかったら 再配置処理へ
}

func initialJoin(hostName string) error {
	//公開鍵を受け取っているはず
	// /root/.ssh/authorised_keyに追記
	// 公開鍵の削除
	// /etc/hostsにipとエイリアスを記述
	return nil
}
