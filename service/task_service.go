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
	task := model.Task{
		Name: name,
		Description: description,
		Reward: reward,
	}
	return s.repository.Create(&task)
}

func (s TaskService) GetAll () ([]model.Task, error) {
	return s.repository.GetAll()
}

func (s TaskService) DeleteByID (taskID int) error {
	return s.repository.Delete(taskID)
}

func (s TaskService) UpdateTaskName(taskID int, newName string) error {
	task, err := s.repository.GetByID(taskID)
	if err != nil {
		return err
	}
	task.Name = newName
	return s.repository.Update(taskID, task)
}

func (s TaskService) UpdateTaskDescription(taskID int, newDescription string) error {
	task, err := s.repository.GetByID(taskID)
	if err != nil {
		return err
	}
	task.Description = newDescription
	return s.repository.Update(taskID, task)
}

func (s TaskService) UpdateTaskReward(taskID, newReward int) error {
	task, err := s.repository.GetByID(taskID)
	if err != nil {
		return err
	}
	task.Reward = newReward
	return s.repository.Update(taskID, task)
}
