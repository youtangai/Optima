package db

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/youtangai/Optima/conductor/config"
)

var (
	DataBase *gorm.DB
)

func init() {
	DataBase = connection()
}

func connection() *gorm.DB {
	dbms := config.DBMS()
	user := config.DBUser()
	pass := config.DBPasswd()
	host := config.DBHost()
	port := config.DBPort()
	dbName := config.DBName()
	connect := user + ":" + pass + "@" + "tcp(" + host + ":" + port + ")/" + dbName
	db, err := gorm.Open(dbms, connect)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
