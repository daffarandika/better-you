package handler

import (
	"betteryou/service"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct{
	taskService	service.TaskService
}

func NewTaskHandler(taskService service.TaskService) *TaskHandler {
	return &TaskHandler{
		taskService: taskService,
	}
}

func (h TaskHandler) CreateNewTask(c echo.Context) {

}
