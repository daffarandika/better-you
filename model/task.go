package model

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Name		string	`gorm:"column:name;type:varchar(100);not null"`
	Description string	`gorm:"column:description;type:varchar(255);not null"`
	Reward 		int		`gorm:"column:reward;type:int;not null"`
}
