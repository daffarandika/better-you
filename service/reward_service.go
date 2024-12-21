package service

import (
	"betteryou/model"
	"betteryou/repository"
)

type RewardService struct {
	repository *repository.RewardRepository
}

func NewRewardService (repository *repository.RewardRepository) *RewardService {
	return &RewardService{
		repository:	repository,
	}
}

func (s RewardService) Create(name, description string, cost int) (*model.Reward, error) {
	reward := &model.Reward{
		Name: name,
		Description: description,
		Cost: cost,
	}
	return reward, s.repository.Create(reward)
}

func (s RewardService) GetAll() ([]model.Reward, error) {
	return s.repository.GetAll()
}

func (s RewardService) GetByID(rewardID int) (model.Reward, error) {
	return s.repository.GetByID(rewardID)
}
