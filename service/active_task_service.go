package service

import (
	"betteryou/model"
	"betteryou/repository"
)

type ActiveTaskService struct{
	activeTaskRepository *repository.ActiveTaskRepository;
	taskRepository *repository.TaskRepository;
}

func NewActiveTaskService(activeTaskRepository *repository.ActiveTaskRepository, taskRepository *repository.TaskRepository) *ActiveTaskService {
	return &ActiveTaskService{ activeTaskRepository: activeTaskRepository, taskRepository: taskRepository }
}

func (s ActiveTaskService) Create (taskID int) error {
	activeTask := &model.ActiveTask{
		TaskID:	taskID,
		Done:	false,
	}
	return s.activeTaskRepository.Create(activeTask)
}

func (s ActiveTaskService) GetAll () ([]model.ActiveTask, error) {
	activeTasks, err := s.activeTaskRepository.GetAll()
	for i, activeTask := range activeTasks {
		task, err := s.taskRepository.GetByID(activeTask.TaskID)
		if err != nil {
			return nil, err
		}
		activeTasks[i].Task = *task
	}
	return activeTasks, err
}

func (s ActiveTaskService) DeleteByID (taskID int) error {
	return s.activeTaskRepository.Delete(taskID)
}

func (s ActiveTaskService) UpdateDoneStatus(activeTaskID int, done bool) error {
	activeTask, err := s.activeTaskRepository.GetByID(activeTaskID)
	if err != nil {
		return err
	}
	activeTask.Done = done
	return s.activeTaskRepository.Update(activeTaskID, activeTask)
}

