package main

import (
	"github.com/youtangai/Optima/conductor/db"
	"github.com/youtangai/Optima/conductor/model"
)

func main() {
	db := db.GetDataBase()
	db.CreateTable(&model.Host{})
}
