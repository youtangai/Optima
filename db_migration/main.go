package main

import (
	"github.com/youtangai/Optima/conductor/db"
	"github.com/youtangai/Optima/conductor/model"
)

func main() {
	db := db.GetDataBase()
	if db.HasTable(&model.LoadIndicator{}) {
		db.DropTable(&model.LoadIndicator{})
	}
	if db.HasTable(&model.Host{}) {
		db.DropTable(&model.Host{})
	}
	db.AutoMigrate(&model.Host{}, &model.LoadIndicator{})
}
