package handler

import (
	"betteryou/service"
	"betteryou/view/component"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct{
	taskService	*service.TaskService
}

func NewTaskHandler(taskService *service.TaskService) *TaskHandler {
	return &TaskHandler{
		taskService: taskService,
	}
}

func (h TaskHandler) CreateNewTask(c echo.Context) error {
	name := c.FormValue("name")
	description := c.FormValue("description")
	reward, err := strconv.Atoi(c.FormValue("reward"))
	if err != nil {
		return err
	}

	task, err := h.taskService.Create(name, description, reward)
	if err != nil {
		return err
	}

	return render(c, component.TaskItem(*task), 200)
}
