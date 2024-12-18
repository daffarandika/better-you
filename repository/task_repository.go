package repository

import (
	"betteryou/model"
	"gorm.io/gorm"
)

type TaskRepository struct{
	db	*gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository{
	return &TaskRepository{db: db}
}

func (r TaskRepository) GetAll() ([]model.Task, error) {
	var tasks []model.Task
	result := r.db.Find(&tasks)
	return tasks, result.Error
}

func (r TaskRepository) GetByID(taskID int) (*model.Task, error) {
	var task model.Task
	result := r.db.First(&task, taskID)
	return &task, result.Error
}

func (r TaskRepository) Create(task *model.Task) error {
	result := r.db.Create(task)
	return result.Error
}

func (r TaskRepository) Update(taskID int, updatedTask *model.Task) error {
	result := r.db.Model(&model.Task{}).Where("id = ?", taskID).Updates(updatedTask)
	return result.Error
}

func (r TaskRepository) Delete(taskID int) error {
	task, err := r.GetByID(taskID)
	if err != nil {
		return err
	}
	result := r.db.Delete(task)
	return result.Error
}
