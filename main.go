package main

import (
	"betteryou/db"
	"betteryou/handler"
	"betteryou/model"
	"betteryou/repository"
	"betteryou/service"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) {
	db.AutoMigrate(&model.Task{})
	db.AutoMigrate(&model.ActiveTask{})
	db.AutoMigrate(&model.Reward{})
	db.AutoMigrate(&model.Transaction{})
}

func main() {
	db, err := db.CreateDb()
	if err != nil {
		log.Fatal("failed to connect to database")
	}
	migrate(db)

	taskRepository := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepository)
	activeTaskRepository := repository.NewActiveTaskRepository(db)
	activeTaskService := service.NewActiveTaskService(activeTaskRepository, taskRepository)

	activeTaskService.Create(1)

	e := echo.New()
	e.Debug = true
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	  Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Static("/static", "static")

	homeHandler := handler.NewHomeHandler(taskService, activeTaskService)

	e.GET("/", homeHandler.HomeGetHandler)

	e.Start(":3000")
}
