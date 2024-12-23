package model

import (
	"gorm.io/gorm"
)

type ActiveTask struct {
	gorm.Model
	TaskID	int		`gorm:"not null"`
	Task	Task	`gorm:"foreignKey:TaskID;references:ID"`
	Done	bool	`gorm:"not null"`
}

func (at *ActiveTask) BeforeCreate(tx *gorm.DB) error {
	var task Task
	result := tx.First(&task, at.TaskID)
	if result.Error != nil {
		return result.Error
	}

	at.Task = task
	return nil
}

func (at *ActiveTask) AfterFind(tx *gorm.DB) error {
	var task Task
	result := tx.First(&task, at.TaskID)
	if result.Error != nil {
		return result.Error
	}

	at.Task = task
	return nil
}

