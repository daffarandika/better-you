package model

type User struct { 
	ID		int		`gorm:"primaryKey"`
	Name	string	`gorm:not null`
	Coins	int		`gorm:"column:coins;type:int;not null"`
}
