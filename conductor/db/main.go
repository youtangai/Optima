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

//GetContainersByHostName is ホスト名からコンテナの一覧を取得する関数
func GetContainersByHostName(hostname string) (*[]model.Container, error) {
	containers := new([]model.Container)
	err := DataBase.Where(&model.Container{Host: hostname}).Find(containers).Error
	if err != nil {
		return containers, err
	}
	return containers, nil
}

//GetContainerByUUID is uuidからコンテナを１つ取得する
func GetContainerByUUID(uuid string) (*model.Container, error) {
	container := new(model.Container)
	for {
		err := DataBase.Where(&model.Container{UUID: uuid}).First(container).Error
		if err != nil {
			return container, err
		}
		if container.Host != "" && container.ContainerID != "" {
			break
		}
	}
	return container, nil
}

//RegistCheckPointDir is チェックポイントパスを登録する関数
func RegistCheckPointDir(chkDirPath, imageName string) error {
	checkpoint := new(model.Checkpoint)
	checkpoint.CheckDir = chkDirPath
	checkpoint.ContainerImage = imageName
	checkpoint.IsRestored = false
	err := DataBase.Create(checkpoint).Error
	if err != nil {
		return err
	}
	return nil
}

//DeleteLoadIndicator is 負荷指標を削除する関数
func DeleteLoadIndicator(hostName string) error {
	loadIndicator := new(model.LoadIndicator)
	err := DataBase.Where(&model.LoadIndicator{HostName: hostName}).First(loadIndicator).Error
	if err != nil {
		return err
	}

	err = DataBase.Delete(&loadIndicator).Error
	if err != nil {
		return err
	}

	return nil
}

//GetCheckPointDirs is レストアされていないチェックポイントの一覧を取得する
func GetCheckPointDirs() (*[]model.Checkpoint, error) {
	checkpoints := new([]model.Checkpoint)
	err := DataBase.Where(&model.Checkpoint{IsRestored: false}).Find(checkpoints).Error
	if err != nil {
		return checkpoints, err
	}
	return checkpoints, err
}

//DeleteCheckPointDir is チェックポイントレコードを削除する
func DeleteCheckPointDir(dir model.Checkpoint) error {
	err := DataBase.Delete(dir).Error
	if err != nil {
		return err
	}
	return nil
}

//GetHostOrderByLoadIndicator is 負荷の高い順にホストを並べて返却する
func GetHostOrderByLoadIndicator() ([]model.LoadIndicator, error) {
	loadIndicators := new([]model.LoadIndicator)
	err := DataBase.Order("load_indicator desc").Find(loadIndicators).Error
	if err != nil {
		return *loadIndicators, err
	}
	return *loadIndicators, nil
}
