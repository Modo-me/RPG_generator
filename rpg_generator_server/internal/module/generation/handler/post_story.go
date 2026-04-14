package handler

import (
	"rpg_generator/internal/module/generation/repository"
	"rpg_generator/internal/module/generation/service"

	"github.com/gin-gonic/gin"
)

type StoryHandler struct {
	storyService *service.StoryService
}

func NewStoryHandler(storyService *service.StoryService) *StoryHandler {
	return &StoryHandler{
		storyService: storyService,
	}
}
func (h *StoryHandler) AddStory(c *gin.Context) {
	var newStory repository.Story
	if err := c.BindJSON(&newStory); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	err := h.storyService.AddStory(&newStory)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to add story"})
		return
	}
	c.IndentedJSON(201, newStory)
}
