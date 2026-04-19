package service

import (
	"rpg_generator/internal/module/task/repository"
)

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
