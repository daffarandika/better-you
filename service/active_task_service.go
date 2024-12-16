package service

// import (
// 	"betteryou/model"
// 	"betteryou/repository"
// )
//
// type ActiveTaskService struct{
// 	activeTaskRepository *repository.ActiveTaskRepository;
// 	taskRepository *repository.TaskRepository;
// }
//
// func NewActiveTaskService(
// 	activeTaskRepository *repository.ActiveTaskRepository,
// 	taskRepository *repository.TaskRepository,
// ) *ActiveTaskService {
// 	return &ActiveTaskService{
// 		activeTaskRepository: activeTaskRepository,
// 		taskRepository: taskRepository,
// 	}
// }
//
// func (s ActiveTaskService) Create (taskID int) error {
// 	searchedTask, err := s.taskRepository.GetByID(taskID)
// 	if err != nil {
// 		return err
// 	}
//
// }
//
// func (s TaskService) GetAll () ([]model.Task, error) {
// 	return s.repository.GetAll()
// }
//
// func (s TaskService) DeleteByID (taskID int) error {
// 	return s.repository.Delete(taskID)
// }
//
// func (s TaskService) UpdateTaskName(taskID int, newName string) error {
// 	task, err := s.repository.GetByID(taskID)
// 	if err != nil {
// 		return err
// 	}
// 	task.Name = newName
// 	return s.repository.Update(taskID, task)
// }
//
// func (s TaskService) UpdateTaskDescription(taskID int, newDescription string) error {
// 	task, err := s.repository.GetByID(taskID)
// 	if err != nil {
// 		return err
// 	}
// 	task.Description = newDescription
// 	return s.repository.Update(taskID, task)
// }
//
// func (s TaskService) UpdateTaskReward(taskID, newReward int) error {
// 	task, err := s.repository.GetByID(taskID)
// 	if err != nil {
// 		return err
// 	}
// 	task.Reward = newReward
// 	return s.repository.Update(taskID, task)
// }
