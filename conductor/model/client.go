package model

type LoadIndicatorJson struct {
	LoadIndicator float64 `json:"load_indicator"`
	HostName      string  `json:"host_name"`
	HostIP        string  `json:"host_ip"`
}
