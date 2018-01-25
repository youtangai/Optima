package controller

import (
	"log"
	"net/http"

	"golang.org/x/net/context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
	"github.com/youtangai/Optima/checkpointer/model"
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
	chkDirPath := "/tmp/" + containerID + "/checkpoints/chk"

	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	err = cli.CheckpointCreate(ctx, containerID, types.CheckpointCreateOptions{CheckpointDir: "/tmp", CheckpointID: "chk"})
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return chkDirPath, nil
}

//sourceDirをtargetIPに送信する関数
func scpCheckpointDir(targetIP, sourceDir string) error {
	log.Printf("targetIP= %s\n", targetIP)
	log.Printf("sourceDir= %s\n", sourceDir)
	return nil
}
