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
	db.AutoMigrate(&model.User{})
}

func main() {
	db, err := db.CreateDb()
	if err != nil {
		log.Fatal("failed to connect to database")
	}

	taskRepository := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepository)
	activeTaskRepository := repository.NewActiveTaskRepository(db)
	activeTaskService := service.NewActiveTaskService(activeTaskRepository, taskRepository)
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	rewardRepository := repository.NewRewardRepository(db)
	rewardService := service.NewRewardService(rewardRepository)
	transactionRepository := repository.NewTransactionRepository(db)
	transactionService := service.NewTransactionService(transactionRepository)


	e := echo.New()
	e.Debug = true
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Static("/static", "static")

	homeHandler := handler.NewHomeHandler(taskService, activeTaskService, userService, rewardService)
	taskHandler := handler.NewTaskHandler(taskService)
	rewardHandler := handler.NewRewardHandler(rewardService, userService, transactionService)
	activeTaskHandler := handler.NewActiveTaskHandler(activeTaskService, userService, transactionService)

	e.GET("/", homeHandler.HomeGetHandler)
	e.POST("/task", taskHandler.CreateNewTask)
	e.POST("/reward", rewardHandler.CreateNewReward)
	e.POST("/redeem-reward/:rewardID", rewardHandler.RedeemReward)
	e.POST("/activate-task/:taskID", activeTaskHandler.CreateActiveTask)
	e.POST("/toggle-active-task/:activeTaskID", activeTaskHandler.ToggleDoneStatus)

	e.Start(":3000")
}
