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

	err := checkpoint(containerID)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	chkDir, err := scpCheckpointDir(containerID)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"chk_dir_path": chkDir,
	})
	return
}

//コンテナをチェックポイントする関数
func checkpoint(containerID string) error {

	log.Printf("containerid = %s\n", containerID)

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

	cmdstr := "docker checkpoint create " + containerID + " " + containerID
	_, err := exec.Command("sh", "-c", cmdstr).Output()
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

//sourceDirをtargetIPに送信する関数
func scpCheckpointDir(containerID string) (string, error) {
	sourceDir := "/var/lib/docker/containers/" + containerID + "/checkpoints/" + containerID
	keyPath := config.GetSecretKeyPath()
	contollerIP := config.GetControllerIP()
	hostName, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	cmdstr := "scp -o StrictHostKeyChecking=no -i " + keyPath + " -r " + sourceDir + " root@" + contollerIP + ":/var/optima/" + hostName + "/"
	output, err := exec.Command("sh", "-c", cmdstr).Output()
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	log.Printf("cmd = %s", cmdstr)
	log.Printf("sourceDir= %s\n", sourceDir)
	log.Printf("output = %s", output)
	chkDir := "/var/optima/" + hostName + "/" + containerID
	return chkDir, nil
}
