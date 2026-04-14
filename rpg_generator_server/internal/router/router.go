package router

import (
	"rpg_generator/internal/module/generation/handler"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Story *handler.StoryHandler
}

func SetUpRouters(h *Handlers) *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/stories", h.Story.AddStory)
	}
	return router
}
