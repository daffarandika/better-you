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
	userService			*service.UserService;
}

func NewHomeHandler(
	taskService			*service.TaskService,
	activeTaskService	*service.ActiveTaskService,
	userService			*service.UserService,
) *HomeHandler {
	return &HomeHandler{
		taskService: taskService,
		activeTaskService: activeTaskService,
		userService: userService,
	}
}

func (h HomeHandler) HomeGetHandler(
	c echo.Context,
) error {
	tasks, err_t := h.taskService.GetAll()
	activeTasks, err_a := h.activeTaskService.GetAll()
	user, err_u := h.userService.GetByID(1)

	if err_t != nil {
		return c.String(http.StatusNotFound, "Error while fetching tasks")
	}

	if err_a != nil {
		return c.String(http.StatusNotFound, "Error while fetching active tasks")
	}

	if err_u != nil {
		return c.String(http.StatusNotFound, "Error while fetching  users")
	}
	return render(c, layout.Home(tasks, activeTasks, user), http.StatusOK)
}
