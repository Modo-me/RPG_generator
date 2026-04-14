package service

import (
	"rpg_generator/internal/module/generation/repository"
)

type StoryService struct {
	storyDao *repository.StoryDao
}

func NewStoryService(storyDao *repository.StoryDao) *StoryService {
	return &StoryService{
		storyDao: storyDao,
	}
}

func (s *StoryService) AddStory(story *repository.Story) error {
	err := s.storyDao.AddStory(story)
	return err
}
