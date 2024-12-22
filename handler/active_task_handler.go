package handler

import (
	"betteryou/service"
	"betteryou/view/component"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ActiveTaskHandler struct {
	activeTaskService	*service.ActiveTaskService
}

func NewActiveTaskHandler (activeTaskService *service.ActiveTaskService) *ActiveTaskHandler {
	return &ActiveTaskHandler{
		activeTaskService:	activeTaskService,
	}
}

func (h ActiveTaskHandler) CreateActiveTask (c echo.Context) error {
	taskID, err := strconv.Atoi(c.Param("taskID"))
	if err != nil {
		return err
	}

	activeTask, err := h.activeTaskService.Create(taskID)
	if err != nil {
		return err
	}

	return render(c, component.ActiveTaskItem(*activeTask), 200)
}

func (h ActiveTaskHandler) ToggleDoneStatus (c echo.Context) error {
	activeTaskID, err := strconv.Atoi(c.Param("activeTaskID"))
	if err != nil {
		return err
	}

	activeTask, err := h.activeTaskService.ToggleDoneStatus(activeTaskID)
	if err != nil {
		return err
	}

	return render(c, component.ActiveTaskItem(*activeTask), 200)
}
