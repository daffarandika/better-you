package model

import "gorm.io/gorm"

type transactionType string

const (
	task_done transactionType = "task_done"
	task_undone transactionType = "task_undone" // if user cancels marking task as done
	item_purcase transactionType =	"item_purcase"
)

type Transaction struct {
	gorm.Model
	TransactionType transactionType `gorm:"column:transaction_type;not null"`
	CoinsAffected	int				`gorm:"column:coins_affected;not null"`
}
