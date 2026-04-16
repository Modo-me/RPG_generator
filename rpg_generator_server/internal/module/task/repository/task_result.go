package repository

import (
	"rpg_generator/internal/module/tools/repository"
	"sync"

	"gorm.io/gorm"
)

type TaskResult struct {
	TaskId    uint                 `gorm:"primaryKey;autoIncrement"`
	Status    string               `gorm:"type:varchar(20)"`
	Dialogues []tool_repo.Dialogue `gorm:"foreignKey:TaskResultId"`
}

type TaskResultDao struct {
	db *gorm.DB
}

var (
	taskResultDao  *TaskResultDao
	taskResultOnce = new(sync.Once)
)

func NewTaskResultDaoInstance(db *gorm.DB) *TaskResultDao {
	taskResultOnce.Do(func() {
		taskResultDao = &TaskResultDao{
			db: db,
		}
	})
	return taskResultDao
}

// CreateTaskResult Create a new TaskResult with status "pending" and return its TaskId
func (dao *TaskResultDao) CreateTaskResult() (uint, error) {
	taskResult := TaskResult{
		Status: "pending",
	}
	err := dao.db.Create(&taskResult).Error
	return taskResult.TaskId, err
}

func (dao *TaskResultDao) GetTaskResultById(id uint) (*TaskResult, error) {
	var taskResult TaskResult
	err := dao.db.Preload("Dialogues").First(&taskResult, id).Error
	return &taskResult, err
}
