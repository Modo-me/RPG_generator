package repository

import (
	"sync"

	"gorm.io/gorm"
)

type Story struct {
	Id        uint   `Gorm:"primaryKey" json:"id"`
	WorldName string `json:"world_name"`
	WorldDesc string `json:"world_desc"`
	Emotion   string `json:"emotion"`
}

type StoryDao struct {
	db *gorm.DB
}

var (
	storyDao  *StoryDao
	storyOnce sync.Once
)

func NewStoryDaoInstance(db *gorm.DB) *StoryDao {
	storyOnce.Do(
		func() {
			storyDao = &StoryDao{
				db: db,
			}
		})
	return storyDao
}

func (dao *StoryDao) AddStory(story *Story) error {
	result := dao.db.Create(story)
	return result.Error
}

func (dao *StoryDao) DeleteStoryById(id uint) error {
	result := dao.db.Delete(&Story{}, id)
	return result.Error
}
