package repository

import (
	"betteryou/model"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	db	*gorm.DB
}

func NewTransactionRepository (db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{
		db:	db,
	}
}

func (r TransactionRepository) Create(transaction *model.Transaction) error {
	result := r.db.Create(transaction)
	return result.Error
}
