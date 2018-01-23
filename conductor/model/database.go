package model

type Host struct {
	ID       int    `sql:AUTO_INCREMENT`
	HostName string `gorm:"not null"`
	HostIP   string `gorm:"not null"`
}

type LoadIndicator struct {
	ID            int `sql:AUTO_INCREMENT`
	HostID        int
	Host          Host `gorm:"ForeignKey:HostID;AssociationForeignKey:ID"`
	LoadIndicator float64
}
