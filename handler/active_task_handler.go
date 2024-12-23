package handler

import (
	"betteryou/model"
	"betteryou/service"
	"betteryou/view/component"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ActiveTaskHandler struct {
	activeTaskService	*service.ActiveTaskService
	userService			*service.UserService
	transactionService	*service.TransactionService
}

func NewActiveTaskHandler (
	activeTaskService	*service.ActiveTaskService,
	userService			*service.UserService,
	transactionService	*service.TransactionService,
) *ActiveTaskHandler {
	return &ActiveTaskHandler{
		activeTaskService:	activeTaskService,
		userService:		userService,
		transactionService:	transactionService,
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

	// update a user's coin
	newCointAmount, err := h.userService.AddCoins(2, activeTask.Task.Reward)
	if err != nil {
		return err
	}

	// create transaction record
	err = h.transactionService.Create(model.TaskDone, activeTask.Task.Reward)
	if err != nil {
		return err
	}

	return render(c, component.ActiveTaskItemWithUserCoinContainer(activeTask, newCointAmount), 200)
}
