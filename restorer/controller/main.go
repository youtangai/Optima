package controller

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/youtangai/Optima/restorer/config"
	"github.com/youtangai/Optima/restorer/model"
)

//RestoreContainerController is レストアエンドポイントの処理
func RestoreContainerController(c *gin.Context) {
	json := new(model.RestoreContainerInfoJSON)
	c.ShouldBindJSON(json)
	containerID := json.ContainerID
	restoreDir := json.RestoreDir

	sourceID, err := scpRestoreDir(restoreDir)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	err = restore(containerID, sourceID)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "restore done",
	})
	return
}

//コンテナをチェックポイントする関数
func restore(restoreTargetContainerID, restoreSourceContainerID string) error {

	log.Printf("targetContainerId = %s\n", restoreTargetContainerID)

	//レストアコマンドを作成
	cmdstr := "docker start --checkpoint=" + restoreSourceContainerID + " --checkpoint-dir=/var/optima/ " + restoreTargetContainerID
	log.Printf("cmd = %s", cmdstr)

	start := time.Now()
	_, err := exec.Command("sh", "-c", cmdstr).Output()
	if err != nil {
		log.Fatal(err)
		return err
	}
	end := time.Now()
	fmt.Printf("restore took time = %v\n", end.Sub(start))

	cmdstr = "rm -rf /var/optima/" + restoreSourceContainerID
	log.Printf("cmd = %s", cmdstr)
	start = time.Now()
	_, err = exec.Command("sh", "-c", cmdstr).Output()
	if err != nil {
		log.Fatal(err)
		return err
	}
	end = time.Now()
	fmt.Printf("remove checkpoint took time = %v\n", end.Sub(start))

	return nil
}

//restoreディレクトリを取得する処理
func scpRestoreDir(restoreDir string) (string, error) {
	keyPath := config.GetSecretKeyPath()
	contollerIP := config.GetControllerIP()

	//レストアフォルダのダウンロード
	cmdstr := "scp -o StrictHostKeyChecking=no -i " + keyPath + " -r root@" + contollerIP + ":" + restoreDir + " /var/optima/"
	start := time.Now()
	output, err := exec.Command("sh", "-c", cmdstr).Output()
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	end := time.Now()
	fmt.Printf("download checkpoint took time = %v\n", end.Sub(start))
	log.Printf("cmd = %s", cmdstr)
	log.Printf("restoreDir= %s\n", restoreDir)
	log.Printf("output = %s", output)
	restoreDirSplit := strings.Split(restoreDir, "/")      //レストアのパスを / で分割
	containerID := restoreDirSplit[len(restoreDirSplit)-1] // 末尾のコンテナIDを取得
	return containerID, nil
}
