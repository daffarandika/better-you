package handler

import (
	"betteryou/service"
	"betteryou/view/layout"
	"net/http"

	"github.com/labstack/echo/v4"
)

type HomeHandler struct {
	taskService	*service.TaskService
}

func NewHomeHandler(s *service.TaskService) *HomeHandler {
	return &HomeHandler{taskService: s}
}

func (h HomeHandler) HomeGetHandler(
	c echo.Context,
) error {
	tasks, err := h.taskService.GetAll()
	if err != nil {
		c.String(http.StatusNotFound, "Error while fetching")
		return render(c, layout.Home(tasks), http.StatusOK)
	}
	return render(c, layout.Home(tasks), http.StatusOK)
}
