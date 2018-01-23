package db

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql driver
	"github.com/youtangai/Optima/conductor/config"
	"github.com/youtangai/Optima/conductor/model"
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

func GetDataBase() *gorm.DB {
	return DataBase
}

func RegistLoadIndicator(json model.LoadIndicatorJson) error {
	hostEntity := new(model.Host)
	err := DataBase.FirstOrCreate(hostEntity, &model.Host{HostName: json.HostName, HostIP: json.HostIP}).Error
	if err != nil {
		return err
	}
	loadIndicatorEntity := new(model.LoadIndicator)
	loadIndicatorEntity.HostID = hostEntity.ID
	loadIndicatorEntity.Host = *hostEntity
	loadIndicatorEntity.LoadIndicator = json.LoadIndicator
	err = DataBase.Create(loadIndicatorEntity).Error
	if err != nil {
		return err
	}
	return nil
}
