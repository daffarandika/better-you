package model

import "gorm.io/gorm"

type TransactionType string

const (
	TaskDone TransactionType = "task_done"
	TaskUndone TransactionType = "task_undone" // if user cancels marking task as done
	ItemPurcase TransactionType =	"item_purcase"
)

type Transaction struct {
	gorm.Model
	TransactionType TransactionType `gorm:"column:transaction_type;not null"`
	CoinsAffected	int				`gorm:"column:coins_affected;not null"`
}
