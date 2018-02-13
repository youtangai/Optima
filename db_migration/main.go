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
	db.AutoMigrate(&model.LoadIndicator{}, &model.Checkpoint{})

	if db.HasTable(&model.Container{}) {
		log.Println("find container table")
	} else {
		log.Println("could not find container table")
		db.AutoMigrate(&model.Container{})
	}

	if db.HasTable(&model.Avatar{}) {
		log.Println("find avatar table")
	} else {
		log.Println("could not find avatar table")
		db.AutoMigrate(&model.Avatar{}).AddForeignKey("uuid", "container(uuid)", "RESTRICT", "RESTRICT")
	}
}

func readContainer() {
	db := db.GetDataBase()
	container := new(model.Container)
	db.First(container)
	log.Printf("container status\n%+v", container)
}
