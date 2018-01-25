package controller

import (
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
	"github.com/youtangai/Optima/checkpointer/config"
	"github.com/youtangai/Optima/checkpointer/model"
)

//CehckpointContainerController is チェックエンドポイントの処理
func CehckpointContainerController(c *gin.Context) {
	json := new(model.CheckpointContainerInfoJSON)
	c.ShouldBindJSON(json)
	containerID := json.ContainerID

	sourceDirPath, err := checkpoint(containerID)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	err = scpCheckpointDir(sourceDirPath)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "checkpoint & scp done",
	})
	return
}

//コンテナをチェックポイントする関数
func checkpoint(containerID string) (string, error) {
	chkID := "chk" //checkpoint dir name

	log.Printf("containerid = %s\n", containerID)
	chkDirPath := "/var/lib/docker/containers/" + containerID + "/checkpoints/" + chkID

	// ctx := context.Background()
	// cli, err := client.NewEnvClient()
	// if err != nil {
	// 	log.Fatal(err)
	// 	return "", err
	// }

	// //chkID という名前で checkpoint作成
	// err = cli.CheckpointCreate(ctx, containerID, types.CheckpointCreateOptions{CheckpointID: chkID})
	// if err != nil {
	// 	log.Fatal(err)
	// 	return "", err
	// }

	cmdstr := "docker checkpoint create " + containerID + " " + chkID
	_, err := exec.Command("sh", "-c", cmdstr).Output()
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return chkDirPath, nil
}

//sourceDirをtargetIPに送信する関数
func scpCheckpointDir(sourceDir string) error {
	keyPath := config.GetSecretKeyPath()
	contollerIP := config.GetControllerIP()
	hostName, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	cmdstr := "scp -i " + keyPath + " -r " + sourceDir + " root@" + contollerIP + ":/var/optima/" + hostName + "/"
	log.Printf("cmd = %s", cmdstr)
	log.Printf("sourceDir= %s\n", sourceDir)
	return nil
}