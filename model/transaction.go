package model

import "gorm.io/gorm"

type transactionType string

const (
	redeem_reward transactionType = "redeem_reward"
	item_purcase transactionType =	"item_purcase"
)

type Transaction struct {
	gorm.Model
	TransactionType transactionType `gorm:"column:transaction_type;not null"`
	CoinsAffected	int				`gorm:"column:coins_affected;not null"`
}
