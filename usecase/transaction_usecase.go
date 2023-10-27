package usecase

import (
	"fmt"
	"go-transaction/model"
	"go-transaction/repository"
)

type TransactionUsecase interface {
	RegisterNewTransaction(bodyRequest model.Transaction) error
	ShowListTransactions() ([]model.TransactionDTO, error)
	ShowHistoryTransactions(customerId string) ([]model.TransactionDTO, error)
	TransactionResponse(transaction []model.Transaction) ([]model.TransactionDTO, error)
}

type transactionUsecase struct {
	repo          repository.TransactionRepository
	bankUsecase   BankUsecase
	custUsecase   CustomerUsecase
	mrchntUsecase MerchantUsecase
}

func (t *transactionUsecase) RegisterNewTransaction(bodyRequest model.Transaction) error {
	err := t.repo.Create(bodyRequest)

	if err != nil {
		return fmt.Errorf("failed to add transaction : %s", err.Error())
	}
	return nil
}

func (t *transactionUsecase) ShowListTransactions() ([]model.TransactionDTO, error) {
	transactions, err := t.repo.List()
	if err != nil {
		return nil, fmt.Errorf("failed to get transaction list : %v", err.Error())
	}

	transactionResponse, err := t.TransactionResponse(transactions)
	if err != nil {
		return nil, err
	}

	return transactionResponse, nil
}

func (t *transactionUsecase) ShowHistoryTransactions(customerId string) ([]model.TransactionDTO, error) {
	transactions, err := t.repo.History(customerId)
	if err != nil {
		return nil, fmt.Errorf("failed to get transaction history : %v", err.Error())
	}

	transactionResponse, err := t.TransactionResponse(transactions)
	if err != nil {
		return nil, err
	}

	return transactionResponse, nil
}

func (t *transactionUsecase) TransactionResponse(transactions []model.Transaction) ([]model.TransactionDTO, error) {
	var transactionResponse []model.TransactionDTO
	for _, transaction := range transactions {
		customer, err := t.custUsecase.GetDetailCustomer(transaction.CustomerId)
		if err != nil {
			return nil, fmt.Errorf("failed to get customer information : %v", err.Error())
		}

		merchant, err := t.mrchntUsecase.GetDetailmerchant(transaction.MerchantId)
		if err != nil {
			return nil, fmt.Errorf("failed to get merchant information : %v", err.Error())
		}

		bank, err := t.bankUsecase.GetDetailBank(transaction.BankId)
		if err != nil {
			return nil, fmt.Errorf("failed to get bank information : %v", err.Error())
		}

		var row model.TransactionDTO
		row.Id = transaction.Id
		row.Customer = customer
		row.Merchant = merchant
		row.Bank = bank
		row.Amount = transaction.Amount
		row.TransactionAt = transaction.TransactionAt

		transactionResponse = append(transactionResponse, row)
	}
	return transactionResponse, nil
}

func NewTransactionUsecase(repo repository.TransactionRepository, bankUsecase BankUsecase, custUsecase CustomerUsecase, mrchntUsecase MerchantUsecase) TransactionUsecase {
	return &transactionUsecase{repo, bankUsecase, custUsecase, mrchntUsecase}
}
