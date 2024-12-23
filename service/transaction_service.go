package service

import (
	"betteryou/model"
	"betteryou/repository"
)

type TransactionService struct {
	transactionRepository	*repository.TransactionRepository
}

func NewTransactionService(transactionRepository *repository.TransactionRepository) *TransactionService {
	return &TransactionService{
		transactionRepository:	transactionRepository,
	}
}

func (s TransactionService) Create(transactionType model.TransactionType, coinsAffected int) error {
	transaction := &model.Transaction{
		TransactionType:	transactionType,
		CoinsAffected:		coinsAffected,
	}
	return s.transactionRepository.Create(transaction)
}
