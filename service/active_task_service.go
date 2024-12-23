package service

import (
	"betteryou/model"
	"betteryou/repository"
	"time"
)

type ActiveTaskService struct{
	activeTaskRepository *repository.ActiveTaskRepository;
	taskRepository *repository.TaskRepository;
}

func NewActiveTaskService(activeTaskRepository *repository.ActiveTaskRepository, taskRepository *repository.TaskRepository) *ActiveTaskService {
	return &ActiveTaskService{ activeTaskRepository: activeTaskRepository, taskRepository: taskRepository }
}

func (s ActiveTaskService) Create (taskID int) (*model.ActiveTask, error) {
	task, err := s.taskRepository.GetByID(taskID)
	if err != nil {
		return nil, err
	}
	activeTask := &model.ActiveTask{
		TaskID:	taskID,
		Task:	*task,
		Done:	false,
	}
	return activeTask, s.activeTaskRepository.Create(activeTask)
}

func (s ActiveTaskService) GetAll () ([]model.ActiveTask, error) {
	activeTasks, err := s.activeTaskRepository.GetAll()
	return activeTasks, err
}

func (s ActiveTaskService) GetAllValid () ([]model.ActiveTask, error) {
	activeTasks, err := s.GetAll()
	if err != nil { 
		return nil, err
	}

	validTasks := []model.ActiveTask{}
	for _, activeTask := range activeTasks {
		taskActivatedLessThanADayAgo := activeTask.UpdatedAt.After(time.Now().Add(-24 * time.Hour))
		taskNotDone := !activeTask.Done

		if taskActivatedLessThanADayAgo || taskNotDone {
			validTasks = append(validTasks, activeTask)
		}
	}

	return validTasks, err
}

func (s ActiveTaskService) DeleteByID (taskID int) error {
	return s.activeTaskRepository.Delete(taskID)
}

func (s ActiveTaskService) ToggleDoneStatus(activeTaskID int) (*model.ActiveTask, error) {
	activeTask, err := s.activeTaskRepository.GetByID(activeTaskID)
	if err != nil {
		return nil, err
	}
	activeTask.Done = !activeTask.Done
	return activeTask, s.activeTaskRepository.Update(activeTaskID, activeTask)
}
