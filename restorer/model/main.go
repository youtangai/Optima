package model

type RestoreContainerInfoJSON struct {
	ContainerID string `json:"container_id"`
	RestoreDir  string `json:"restore_dir"`
}
