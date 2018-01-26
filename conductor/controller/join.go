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

//CreateDirController is ディレクトリを作成するコントローラ
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

//JoinController is 参加を受け付けるコントローラ
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
	log.Println("relocation start")
	//TODO再配置処理の実装
	//zun.checkpointテーブルの確認
	checkpoints, err := db.GetCheckPointDirs()
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	//チェックポイントが0でなければ
	if len(*checkpoints) != 0 {
		//チェックポイントのレストア試行
		for _, checkpoint := range *checkpoints {
			uuid, err := createContainer(checkpoint.ContainerImage)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				log.Fatal(err)
			}
			targetContainer, err := db.GetContainerByUUID(uuid)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				log.Fatal(err)
			}
			err = restoreContainer(targetContainer.ContainerID, checkpoint.CheckDir, targetContainer.Host)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				log.Fatal(err)
			}
		}
	}
	//負荷順にサーバを取得
	hosts, err := db.GetHostOrderByLoadIndicator()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Fatal(err)
	}
	for _, host := range hosts {
		//そのホスト内のコンテナすべてを取得
		containers, err := db.GetContainersByHostName(host.HostName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			log.Fatal(err)
		}
		//コンテナの数が2未満であればこのホストはスキップする
		if len(*containers) < 2 {
			continue
		}
		//コンテナを回す
		for _, container := range *containers {
			//コンテナを作る
			uuid, err := createContainer(container.Image)
			if err != nil {
				log.Println("failed create container")
				log.Println(err)
				continue
			}
			//元のコンテナをチェックポイントする
			chkDir, err := checkpointContainer(container.ContainerID, container.Host)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				log.Fatal(err)
			}
			//uuidからターゲットコンテナを特定
			targetContainer, err := db.GetContainerByUUID(uuid)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				log.Fatal(err)
			}
			//レストアする
			err = restoreContainer(targetContainer.ContainerID, chkDir, targetContainer.Host)
			if err != nil {
				log.Println("failed restore container")
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				log.Fatal(err)
			}
			//元のコンテナを削除する
			err = deleteContainer(container.UUID)
			if err != nil {
				log.Println("failed restore container")
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				log.Fatal(err)
			}
		}
	}

}

//一番最初の参加時の処理
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
