package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *TaskHandler) QueryTaskResult(c *gin.Context) {
	taskIdStr := c.Param("Id")
	taskId64, err := strconv.ParseUint(taskIdStr, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid taskId"})
		return
	}

	taskId := uint(taskId64)
	taskStatus, err := h.taskService.QueryTaskResult(taskId)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to query task result"})
		return
	}
	c.JSON(200, gin.H{"taskId": taskId, "status": taskStatus, "result": "task result data"})
}
