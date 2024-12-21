package handler

import (
	"betteryou/service"
	"betteryou/view/component"
	"strconv"
	"github.com/labstack/echo/v4"
)

type RewardHander struct {
	rewardService	*service.RewardService
}

func NewRewardHandler(rewardService	*service.RewardService) *RewardHander {
	return &RewardHander{
		rewardService:	rewardService,
	}
}

func (h RewardHander) CreateNewReward(c echo.Context) error {
	name := c.FormValue("name")
	description := c.FormValue("description")
	cost, err := strconv.Atoi(c.FormValue("cost"))
	if err != nil {
		return err
	}

	reward, err := h.rewardService.Create(name, description, cost)
	if err != nil {
		return err
	}

	return render(c, component.RewardItem(*reward), 200)
}
