package handler

import (
	"betteryou/service"
	"betteryou/view/layout"
	"net/http"

	"github.com/labstack/echo/v4"
)

type HomeHandler struct {
	taskService			*service.TaskService;
	activeTaskService	*service.ActiveTaskService;
}

func NewHomeHandler(taskService *service.TaskService, activeTaskService *service.ActiveTaskService) *HomeHandler {
	return &HomeHandler{taskService: taskService, activeTaskService: activeTaskService}
}

func (h HomeHandler) HomeGetHandler(
	c echo.Context,
) error {
	tasks, err_t := h.taskService.GetAll()
	activeTasks, err_a := h.activeTaskService.GetAll()

	if err_t != nil {
		return c.String(http.StatusNotFound, "Error while fetching")
	}

	if err_a != nil {
		return c.String(http.StatusNotFound, "Error while fetching")
	}
	return render(c, layout.Home(tasks, activeTasks), http.StatusOK)
}
