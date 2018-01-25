package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/youtangai/Optima/checkpointer/model"

	_ "github.com/docker/docker/api/types"
	_ "github.com/docker/docker/api/types/container"
	_ "github.com/docker/docker/client"
	_ "golang.org/x/net/context"
)

//CehckpointContainerController is チェックエンドポイントの処理
func CehckpointContainerController(c *gin.Context) {
	json := new(model.CheckpointContainerInfoJSON)
	c.ShouldBindJSON(json)
	containerID := json.ContainerID
	targetIP := json.TargetIP

	sourceDirPath, err := checkpoint(containerID)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	err = scpCheckpointDir(targetIP, sourceDirPath)
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
	log.Printf("containerid = %s\n", containerID)
	chkDirPath := "/tmp/oiah3th4ihaoihgoiehoac/checkpoints/chk1"
	return chkDirPath, nil
}

//sourceDirをtargetIPに送信する関数
func scpCheckpointDir(targetIP, sourceDir string) error {
	log.Printf("targetIP= %s\n", targetIP)
	log.Printf("sourceDir= %s\n", sourceDir)
	return nil
}
