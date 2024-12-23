package model

import (
	"fmt"

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
	fmt.Println(at.Task)
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

func (at *ActiveTask) AfterUpdate(tx *gorm.DB) error {
	var user *User
	result := tx.First(&user, 2)
	if result.Error != nil {
		return result.Error
	}

	var task Task
	result = tx.First(&task, at.TaskID)
	if result.Error != nil {
		return result.Error
	}

	at.Task = task

	user.Coins += task.Reward
	result = tx.Model(&User{}).Where("id = ?", 2).Updates(user)
	return result.Error
}
