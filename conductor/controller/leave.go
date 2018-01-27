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
	log.Printf("leave: hostname = %s", hostName)
	//受け取ったホスト名のコンテナを調べる
	containers, err := db.GetContainersByHostName(hostName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		log.Fatal(err)
	}
	//コンテナの数が0なら終わり
	if len(*containers) == 0 {
		log.Println("leave:delete load_indicator")
		err = db.DeleteLoadIndicator(hostName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			log.Fatal(err)
		}
		log.Println("leave:delete load_indicator done")
		c.JSON(http.StatusOK, gin.H{"message": "no container need checkpoint and restore"})
		return
	}
	//コンテナを１つ選択
	for _, container := range *containers {
		log.Printf("leave:current container info = %+v", container)
		//同イメージでコンテナ作成を依頼 このとき uuidを取得
		uuid, err := createContainer(container.Image)
		log.Printf("leave:created container's uuid = %s", uuid)
		if err != nil { //コンテナの作成ができなかったら
			log.Println("leave:cannot create container")
			//チェックポイントする
			chkDirPath, err := checkpointContainer(container.ContainerID, container.Host)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err})
				log.Fatal(err)
			}
			log.Printf("leave:checkpointdirpath = %s", chkDirPath)
			//chkdirpathをDBへ登録
			err = db.RegistCheckPointDir(chkDirPath, container.Image)
			if err != nil {
				log.Println("leave:cannot regist chkdir into db")
			}
			//コンテナを削除する
			err = deleteContainer(container.UUID)
			if err != nil {
				log.Println("cannot delete container")
				c.JSON(http.StatusInternalServerError, gin.H{"error": err})
				log.Fatal(err)
			}
			log.Println("leave:checkpoint container " + container.ContainerID)
		} else {
			log.Println("leave:created restore container")
			//レストアするコンテナを取得
			targetContainer, err := db.GetContainerByUUID(uuid)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err})
				log.Fatal(err)
			}
			log.Printf("leave:restore target container info = %+v", *targetContainer)
			//チェックポイントする
			chkdirpath, err := checkpointContainer(container.ContainerID, container.Host)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err})
				log.Fatal(err)
			}
			log.Printf("leave:chkdirpaht = %s", chkdirpath)
			//レストアする
			err = restoreContainer(targetContainer.ContainerID, chkdirpath, targetContainer.Host)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err})
				log.Fatal(err)
			}
			log.Println("leave:restored container " + container.ContainerID)
		}
	}
	//load_indicator削除
	log.Println("leave:delete load_indicator")
	err = db.DeleteLoadIndicator(hostName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		log.Fatal(err)
	}
	log.Println("leave:delete load_indicator done")
	c.JSON(http.StatusOK, gin.H{"message": "leave_suceed"})
	return
}
