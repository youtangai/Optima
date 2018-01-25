package model

type RestoreContainerInfoJSON struct {
	ContainerID    string `json:"container_id"`
	RestoreDirPath string `json:"restore_dir_path"`
}
