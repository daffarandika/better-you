package repository

import (
	"betteryou/model"
	"gorm.io/gorm"
)

type ActiveTaskRepository struct{
	db	*gorm.DB
}

func NewActiveTaskRepository(db *gorm.DB) *ActiveTaskRepository{
	return &ActiveTaskRepository{db: db}
}

func (r ActiveTaskRepository) GetAll() ([]model.ActiveTask, error) {
	var activeTasks []model.ActiveTask
	result := r.db.Find(&activeTasks)
	return activeTasks, result.Error
}

func (r ActiveTaskRepository) GetByID(activeTaskID int) (*model.ActiveTask, error) {
	var activeTask model.ActiveTask
	result := r.db.First(&activeTask, activeTaskID)
	return &activeTask, result.Error
}

func (r ActiveTaskRepository) Create(activeTask *model.ActiveTask) error {
	result := r.db.Create(activeTask)
	return result.Error
}

func (r ActiveTaskRepository) Update(activeTaskID int, updatedActiveTask *model.ActiveTask) error {
	result := r.db.Model(&model.Task{}).Where("id = ?", activeTaskID).Updates(updatedActiveTask)
	return result.Error
}

func (r ActiveTaskRepository) Delete(activeTaskID int) error {
	activeTask, err := r.GetByID(activeTaskID)
	if err != nil {
		return err
	}
	result := r.db.Delete(activeTask)
	return result.Error
}
