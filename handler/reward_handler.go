package handler

import (
	"betteryou/service"
	"betteryou/view/component"
	"strconv"
	"github.com/labstack/echo/v4"
)

type RewardHandler struct {
	rewardService		*service.RewardService
	userService			*service.UserService
	transactionService	*service.TransactionService
}

func NewRewardHandler(
	rewardService	*service.RewardService,
	userService		*service.UserService,
	transactionService	*service.TransactionService,
) *RewardHandler {
	return &RewardHandler{
		rewardService:		rewardService,
		userService:		userService,
		transactionService:	transactionService,
	}
}

func (h RewardHandler) CreateNewReward(c echo.Context) error {
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

func (h RewardHandler) RedeemReward(c echo.Context) error {
	rewardID, err := strconv.Atoi(c.Param("rewardID"))
	if err != nil {
		return err
	}

	reward, err := h.rewardService.GetByID(rewardID)
	if err != nil {
		return err
	}

	userCoinAmount, err := h.userService.AddCoins(2, -reward.Cost)

	return render(c, component.UserCoinContainer(userCoinAmount), 200)

}
