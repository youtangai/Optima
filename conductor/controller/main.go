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
	log.Println("pub key not found")
	//TODO再配置処理の実装
	//zun.checkpointテーブルの確認 及びレストア試行
	//高負荷サーバチェック 及びレストア試行
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

//LeaveController is 脱退処理のコントローラ
func LeaveController(c *gin.Context) {
	json := new(model.LeaveJson)
	c.ShouldBindJSON(json)
	hostName := json.HostName
	//受け取ったホスト名のコンテナを調べる
	containers, err := db.GetContainersByHostName(hostName)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	//コンテナの数が0なら終わり
	if len(*containers) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "no container need checkpoint&restore"})
		return
	}
	//コンテナを１つ選択
	for _, container := range *containers {
		//同イメージでコンテナ作成を依頼 このとき uuidを取得
		uuid, err := createContainer(container.Image)
		if err != nil { //コンテナの作成ができなかったら
			log.Println("leave:cannot create container")
			//チェックポイントする
			chkDirPath, err := checkpointContainer(container.ContainerID, container.Host)
			if err != nil {
				log.Fatal(err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": err})
				return
			}
			//chkdirpathをDBへ登録
			db.RegistCheckPointDir(chkDirPath, container.Image)
			//コンテナを削除する
			err = deleteContainer(container.UUID)
			if err != nil {
				log.Fatal(err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": err})
				return
			}
			log.Println("leave:checkpoint container " + container.ContainerID)
		}
		log.Println("leave:created restore container")
		//レストアするコンテナを取得
		targetContainer, err := db.GetContainerByUUID(uuid)
		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		//チェックポイントする
		chkdirpath, err := checkpointContainer(container.ContainerID, container.Host)
		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		//レストアする
		err = restoreContainer(targetContainer.ContainerID, chkdirpath, targetContainer.Host)
		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		log.Println("leave:restored container " + container.ContainerID)
	}
	//load_indicator削除
	log.Println("leave:delete load_indicator")
	err = db.DeleteLoadIndicator(hostName)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	log.Println("leave:delete load_indicator done")
	c.JSON(http.StatusOK, gin.H{"message": "leave_suceed"})
	return
}

func createContainer(imageName string) (string, error) {
	//openstackにコンテナ作成を依頼
	return "uuid", nil
}

func checkpointContainer(containerID, hostName string) (string, error) {
	return "/var/optima/hostname/containerid", nil
}

func deleteContainer(uuid string) error {
	//delete container
	return nil
}

func restoreContainer(containerID, restoreDir, hostName string) error {
	return nil
}
