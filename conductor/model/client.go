package model

type LoadIndicatorJson struct {
	LoadIndicator float64 `json:"load_indicator" binding:"required"`
	HostName      string  `json:"host_name" binding:"required"`
	HostIP        string  `json:"host_ip" binding:"required"`
}

type JoinJson struct {
	HostName string `json:"host_name" binding:"required"`
}
