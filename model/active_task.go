package model

import "gorm.io/gorm"

type ActiveTask struct {
	gorm.Model
	TaskID	int		`gorm:"not null"`
	Task	Task	`gorm:"foreignKey:TaskID;references:ID"`
	Done	bool	`gorm:"not null"`
}
