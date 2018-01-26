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

const (
	DBMS = "mysql"
)

func init() {
	DataBase = connection()
	DataBase.SingularTable(true)
}

func connection() *gorm.DB {
	user := config.DBUser()
	pass := config.DBPasswd()
	host := config.DBHost()
	port := config.DBPort()
	dbName := config.DBName()
	connectionString := user + ":" + pass + "@" + "tcp(" + host + ":" + port + ")/" + dbName + "?parseTime=true"
	db, err := gorm.Open(DBMS, connectionString)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func GetDataBase() *gorm.DB {
	return DataBase
}

//RegistLoadIndicator is 負荷指標登録処理
func RegistLoadIndicator(json model.LoadIndicatorJson) error {
	hostEntity := new(model.LoadIndicator)

	if DataBase.Where(&model.LoadIndicator{HostName: json.HostName, HostIP: json.HostIP}).First(&hostEntity).RecordNotFound() {
		//見つからなかった時
		hostEntity.HostIP = json.HostIP
		hostEntity.HostName = json.HostName
		hostEntity.LoadIndicator = json.LoadIndicator
		err := DataBase.Create(hostEntity).Error
		if err != nil {
			log.Fatal(err)
			return err
		}
		return nil
	}

	//見つかった時
	hostEntity.LoadIndicator = json.LoadIndicator
	err := DataBase.Save(hostEntity).Error
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

//GetIPAddrByHostName is ホスト名で検索しIPアドレスを返却する関数
func GetIPAddrByHostName(hostname string) (string, error) {
	hostEntity := new(model.LoadIndicator)
	err := DataBase.Where(&model.LoadIndicator{HostName: hostname}).First(hostEntity).Error
	if err != nil {
		return "", err
	}
	IPAddr := hostEntity.HostIP
	return IPAddr, nil
}
