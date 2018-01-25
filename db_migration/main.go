package main

import (
	"log"

	"github.com/youtangai/Optima/conductor/db"
	"github.com/youtangai/Optima/conductor/model"
)

func main() {
	log.Println("db check start")
	db := db.GetDataBase()
	db.SingularTable(true)
	db.LogMode(true)
	if db.HasTable(&model.LoadIndicator{}) {
		log.Println("load indicator is exist")
	}
	if db.HasTable(&model.Host{}) {
		log.Println("hosts is exist")
	}
	db.AutoMigrate(&model.Host{}, &model.LoadIndicator{})

	if db.HasTable(&model.Container{}) {
		log.Println("find container table")
	} else {
		log.Println("could not find container table")
		db.AutoMigrate(&model.Container{})
	}
}

func readContainer() {
	db := db.GetDataBase()
	container := new(model.Container)
	db.First(container)
	log.Printf("container status\n%v", container)
}
