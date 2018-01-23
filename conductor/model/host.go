package model

type Host struct {
	ID       int `sql:AUTO_INCREMENT`
	HostName string
	HostIP   string
}
