package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/youtangai/Optima/checkpointer/model"
)

func CehckpointContainerController(c *gin.Context) {
	json := new(model.CheckpointContainerInfoJSON)
	c.ShouldBindJSON(json)
	containerID := json.ContainerID
	targetIP := json.TargetIP

	sourceDirPath := checkpoint(containerID)
	err := scpCheckpointDir(targetIP, sourceDirPath)
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
}

func checkpoint(containerID string) string {
	log.Printf("container id=%s\n", containerID)
	chkDirPath := "/tmp/oiah3th4ihaoihgoiehoac/checkpoints/chk1"
	return chkDirPath
}

func scpCheckpointDir(targetIP, sourceDir string) error {
	log.Printf("targetIP =%s\n", targetIP)
	log.Printf("sourceDir =%s\n", sourceDir)
	return nil
}
