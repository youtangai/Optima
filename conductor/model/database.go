package model

import "time"

type LoadIndicator struct {
	ID            int    `gorm:"primary_key;AUTO_INCREMENT"`
	HostName      string `gorm:"not null"`
	HostIP        string `gorm:"not null"`
	LoadIndicator float64
}

type Container struct {
	CreatedAt       time.Time `gorm:"column:created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at"`
	ID              int       `gorm:"column:id;not null;primary_key;AUTO_INCREMENT"`
	ProjcetID       string    `gorm:"column:project_id"`
	UserID          string    `gorm:"column:user_id"`
	UUID            string    `gorm:"column:uuid;type:varchar(36);unique"`
	Name            string    `gorm:"column:name"`
	Image           string    `gorm:"column:image"`
	Command         string    `gorm:"column:command"`
	Status          string    `gorm:"column:status;type:varchar(20)"`
	Environment     string    `gorm:"column:environment;type:text"`
	ContainerID     string    `gorm:"column:container_id"`
	Memory          string    `gorm:"column:memory"`
	TaskState       string    `gorm:"column:task_state;type:varchar(20)"`
	CPU             float32   `gorm:"column:cpu;type:float"`
	WorkDir         string    `gorm:"column:workdir"`
	Ports           string    `gorm:"column:ports;type:text"`
	HostName        string    `gorm:"column:hostname"`
	Labels          string    `gorm:"column:labels;type:text"`
	StatusReason    string    `gorm:"column:status_reason;type:text"`
	ImagePullPolicy string    `gorm:"column:image_pull_policy;type:text"`
	Meta            string    `gorm:"column:meta;type:text"`
	Addresses       string    `gorm:"column:addresses;type:text"`
	Host            string    `gorm:"column:host"`
	RestartPolicy   string    `gorm:"column:restart_policy"`
	StatusDetail    string    `gorm:"column:status_detail;type:varchar(50)"`
	ImageDriver     string    `gorm:"column:image_driver"`
	Interactive     bool      `gorm:"column:interactive"`
	WebsocketURL    string    `gorm:"column:websocket_url"`
	WebsocketToken  string    `gorm:"column:websocket_token"`
	SecurityGroups  string    `gorm:"column:security_groups;type:text"`
	AutoRemove      bool      `gorm:"column:auto_remove"`
	Avatar          []Avatar
}

type Checkpoint struct {
	ID             int    `gorm:"primary_key;AUTO_INCREMENT"`
	ContainerImage string `gorm:"not null"`
	CheckDir       string `gorm:"not null"`
	IsRestored     bool   `gorm:"not null;default:false"`
}

type Avatar struct {
	UUID        string
	Host        string
	ContainerID string
	CreatedAt   time.Time
	Continer    Container `gorm:"ForeignKey:UUID;AssociationForeignKey:UUID"`
}
