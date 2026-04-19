package service

import "rpg_generator/internal/module/task/repository"

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
