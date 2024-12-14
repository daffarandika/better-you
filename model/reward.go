package model

import "gorm.io/gorm"

type Reward struct {
	gorm.Model
	Name		string	`gorm:"column:name;type:varchar(100);not null"`
	Description string	`gorm:"column:description;type:varchar(255);not null"`
	Cost 		int		`gorm:"column:cost;type:int;not null"`
}
