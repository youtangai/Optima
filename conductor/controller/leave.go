package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/youtangai/Optima/conductor/db"
	"github.com/youtangai/Optima/conductor/model"
)

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
