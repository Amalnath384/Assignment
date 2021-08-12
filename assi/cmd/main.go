package main

import (
	"project/assi/pkg/model"
	db "project/assi/pkg/repository"
	."project/assi/router"
)

func main() {
	dbHost := "127.0.0.1:27017"
	db.Init(&model.Database{
		Driver:   "mongodb",
		Endpoint: dbHost})
	defer db.Exit()

	Router()

}
