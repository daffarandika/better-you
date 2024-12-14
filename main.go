package main

import (
	"betteryou/db"
	"betteryou/model"
	"log"
)

func main() {
	db, err := db.CreateDb()
	if err != nil {
		log.Fatal("failed to connect to database")
	}
	db.AutoMigrate(&model.Task{})
	db.AutoMigrate(&model.ActiveTask{})
	db.AutoMigrate(&model.Reward{})
	db.AutoMigrate(&model.Transaction{})
}
