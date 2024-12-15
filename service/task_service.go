package service

import (
	"betteryou/model"
	"betteryou/repository"
)

type TaskService struct{
	repository *repository.TaskRepository
}

func NewTaskService(repository *repository.TaskRepository) *TaskService {
	return &TaskService{repository: repository}
}

func (s TaskService) Create (name, description string, reward int) error {
	task := model.Task{Name: name, Description: description, Reward: reward}
	return s.repository.Create(&task)
}

func (s TaskService) GetAll () ([]model.Task, error) {
	return s.repository.GetAll()
}

func (s TaskService) Delete (taskID int) error {
	return s.repository.Delete(taskID)
}
