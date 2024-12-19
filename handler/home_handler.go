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
	rewardService		*service.RewardService;
}

func NewHomeHandler(
	taskService			*service.TaskService,
	activeTaskService	*service.ActiveTaskService,
	userService			*service.UserService,
	rewardService		*service.RewardService,
) *HomeHandler {
	return &HomeHandler{
		taskService:		taskService,
		activeTaskService:	activeTaskService,
		userService:		userService,
		rewardService:		rewardService,
	}
}

func (h HomeHandler) HomeGetHandler(
	c echo.Context,
) error {
	tasks, err := h.taskService.GetAll()
	if err != nil {
		return err
	}

	activeTasks, err := h.activeTaskService.GetAll()
	if err != nil {
		return err
	}

	user, err := h.userService.GetByID(1)
	if err != nil {
		return err
	}

	rewards, err := h.rewardService.GetAll()
	if err != nil {
		return err
	}
	return render(c, layout.Home(tasks, activeTasks, rewards, user), http.StatusOK)
}
