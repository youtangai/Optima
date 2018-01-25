package controller

import (
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/youtangai/Optima/checkpointer/config"
	"github.com/youtangai/Optima/checkpointer/model"
)

//RestoreContainerController is レストアエンドポイントの処理
func RestoreContainerController(c *gin.Context) {
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
func restore(containerID, restoreDirPath string) error {

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

//restoreディレクトリを取得する処理
func scpRestoreDir(restoreDir string) (string, error) {
	keyPath := config.GetSecretKeyPath()
	contollerIP := config.GetControllerIP()
	hostName, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
		return "", nil
	}
	cmdstr := "scp -o StrictHostKeyChecking=no -i " + keyPath + " -r root@" + contollerIP + ":" + restoreDir + " /var/optima/"
	output, err := exec.Command("sh", "-c", cmdstr).Output()
	log.Printf("cmd = %s", cmdstr)
	log.Printf("restoreDir= %s\n", restoreDir)
	log.Printf("output = %s", output)
	restoreDirSplit := strings.Split(restoreDir, "/")      //レストアのパスを / で分割
	containerID := restoreDirSplit[len(restoreDirSplit)-1] // 末尾のコンテナIDを取得
	restorePath := "/var/optima/" + containerID            //レストアするディレクトリのパスを生成
	return restorePath, nil
}
