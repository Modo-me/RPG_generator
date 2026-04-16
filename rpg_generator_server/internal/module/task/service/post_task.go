package service

import (
	"rpg_generator/internal/module/task/repository"
)

type TaskService struct {
	TaskDao       *repository.TaskDao
	TaskResultDao *repository.TaskResultDao
}

func NewTaskService(TaskDao *repository.TaskDao, TaskResultDao *repository.TaskResultDao) *TaskService {
	return &TaskService{
		TaskDao:       TaskDao,
		TaskResultDao: TaskResultDao,
	}
}

// CreateTask Todo: add logic to process task with agent
func (s *TaskService) CreateTask(task *repository.Task) (uint, error) {
	//add task to database
	err := s.TaskDao.AddTask(task)
	if err != nil {
		return 0, err
	}
	//create task result for response
	taskId, err := s.TaskResultDao.CreateTaskResult()
	if err != nil {
		return 0, err
	}

	return taskId, nil
}

// runTask Todo: add logic to process task with agent
func (s *TaskService) runTask(task *repository.Task) error {
	return nil
}
