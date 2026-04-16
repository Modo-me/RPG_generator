package handler

import (
	"rpg_generator/internal/module/task/repository"
	"rpg_generator/internal/module/task/service"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	taskService *service.TaskService
}

func NewTaskHandler(taskService *service.TaskService) *TaskHandler {
	return &TaskHandler{
		taskService: taskService,
	}
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var newTask repository.Task
	if err := c.BindJSON(&newTask); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	taskId, err := h.taskService.CreateTask(&newTask)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to add task"})
		return
	}
	c.JSON(201, gin.H{"taskId": taskId, "status": "pending"})
}
