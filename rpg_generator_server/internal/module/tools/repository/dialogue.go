package tool_repo

import (
	"sync"

	"gorm.io/gorm"
)

type Dialogue struct {
	Id           int64    `gorm:"primaryKey;autoIncrement"`
	Chat         []string `gorm:"type:text[]"`
	TaskResultId int64    `gorm:"index"`
}

type DialogueDao struct {
	db *gorm.DB
}

var (
	dialogueDao  *DialogueDao
	dialogueOnce = new(sync.Once)
)

func NewDialogueDaoInstance(db *gorm.DB) *DialogueDao {
	dialogueOnce.Do(func() {
		dialogueDao = &DialogueDao{
			db: db,
		}
	})
	return dialogueDao
}

func (dao *DialogueDao) AddDialogue(dialogue *Dialogue) error {
	return dao.db.Create(dialogue).Error
}
