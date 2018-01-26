package controller

import (
	"log"
	"net/http"
	"os"
	"os/exec"

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
		return
	}
	hostName := json.HostName
	err = os.Chdir("/var")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = os.MkdirAll("optima/"+hostName, 0777)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"dir_path": "/var/optima/" + hostName})
	return
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
	log.Println("pub key not found")
}

func initialJoin(hostName string) error {
	log.Println("pub key is exists")
	// /root/.ssh/authorised_keyに公開鍵を追記
	cmdstr := "cat /var/optima/" + hostName + "/" + PublicKeyName + " >> /root/.ssh/authorized_keys"
	log.Printf("cmdstr = %s", cmdstr)
	_, err := exec.Command("sh", "-c", cmdstr).Output()
	if err != nil {
		log.Fatal(err)
		return err
	}
	// 公開鍵の削除
	cmdstr = "rm -f /var/optima/" + hostName + "/" + PublicKeyName
	_, err = exec.Command("sh", "-c", cmdstr).Output()
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Println("delete pub key")
	// /etc/hostsにipとエイリアスを記述
	err = os.Chdir("/etc") // etcへ移動
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Println("cd /etc")

	//ipaddrを取得
	ipAddr, err := db.GetIPAddrByHostName(hostName)
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Printf("ip addr = %s", ipAddr)

	//入力したい文字列を生成
	inputString := ipAddr + " " + hostName + "\n"
	log.Printf("inputString = %s", inputString)
	file, err := os.OpenFile("hosts", os.O_WRONLY|os.O_APPEND, 0666) //ファイルを開く
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer file.Close()

	//書込みe
	file.Write(([]byte)(inputString))
	return nil
}
