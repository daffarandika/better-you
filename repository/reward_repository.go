package repository

import (
	"betteryou/model"
	"gorm.io/gorm"
)

type RewardRepository struct {
	db	*gorm.DB
}

func NewRewardRepository(db *gorm.DB) *RewardRepository {
	return &RewardRepository {
		db:	db,
	}
}

func (r RewardRepository) GetAll() ([]model.Reward, error) {
	var rewards []model.Reward
	result := r.db.Find(&rewards)
	return rewards, result.Error
}

func (r RewardRepository) GetByID(rewardID int) (model.Reward, error) {
	var reward model.Reward
	result := r.db.First(&reward, rewardID)
	return reward, result.Error
}

func (r RewardRepository) Create(reward *model.Reward) error {
	result := r.db.Create(reward)
	return result.Error
}

func (r RewardRepository) Update(rewardID int, reward *model.Reward) error {
	result := r.db.Model(&model.Reward{}).Where("id = ?", rewardID).Updates(reward)
	return result.Error
}

func (r RewardRepository) Delete(rewardID int) error {
	reward, err := r.GetByID(rewardID)
	if err != nil {
		return err
	}
	result := r.db.Delete(reward)
	return result.Error
}
