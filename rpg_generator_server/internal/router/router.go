package router

import (
	"rpg_generator/internal/module/task/handler"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Task *handler.TaskHandler
}

func SetUpRouters(h *Handlers) *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/tasks", h.Task.CreateTask)
		api.GET("/tasks/:id", h.Task.QueryTaskResult)
	}
	return router
}
