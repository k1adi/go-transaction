package usecase

import (
	"fmt"
	"go-transaction/model"
	"go-transaction/repository"
)

type TransactionUsecase interface {
	RegisterNewTransaction(bodyRequest model.Transaction) error
}

type transactionUsecase struct {
	repo repository.TransactionRepository
}

func (t *transactionUsecase) RegisterNewTransaction(bodyRequest model.Transaction) error {
	err := t.repo.Create(bodyRequest)

	if err != nil {
		return fmt.Errorf("failed to add transaction : %s", err.Error())
	}
	return nil
}

func NewTransactionUsecase(repo repository.TransactionRepository) TransactionUsecase {
	return &transactionUsecase{repo}
}
