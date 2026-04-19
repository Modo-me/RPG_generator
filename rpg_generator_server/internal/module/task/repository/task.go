package repository

import (
	"sync"

	"gorm.io/gorm"
)

type Task struct {
	Id        int64  `Gorm:"primaryKey" json:"id"`
	WorldName string `json:"world_name"`
	WorldDesc string `json:"world_desc"`
	Emotion   string `json:"emotion"`
}

type TaskDao struct {
	db *gorm.DB
}

var (
	taskDao  *TaskDao
	taskOnce sync.Once
)

func NewTaskDaoInstance(db *gorm.DB) *TaskDao {
	taskOnce.Do(
		func() {
			taskDao = &TaskDao{
				db: db,
			}
		})
	return taskDao
}

func (dao *TaskDao) AddTask(task *Task) error {
	result := dao.db.Create(task)
	return result.Error
}

func (dao *TaskDao) DeleteTaskByIdw(id uint) error {
	result := dao.db.Delete(&Task{}, id)
	return result.Error
}
