package model

import "gorm.io/gorm"

type ActiveTask struct {
	gorm.Model
	TaskID	int	
	Task	Task	`gorm:"foreignKey:TaskID;references:ID"`
	Done	bool
}
