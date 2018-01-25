package model

type ContainerCheckpointInfoJSON struct {
	ContainerID string `json:"container_id"`
	TargetIP    string `json:"target_ip"`
}
