package repository

import (
	"database/sql"
	"go-transaction/model"
	"go-transaction/utils/constant"
)

type TransactionRepository interface {
	Create(bodyRequest model.Transaction) error
	List() ([]model.Transaction, error)
	History(customerId string) ([]model.Transaction, error)
	Scan(rows *sql.Rows) ([]model.Transaction, error)
}

type transactionRepository struct {
	db *sql.DB
}

func (t *transactionRepository) Create(bodyRequest model.Transaction) error {
	if _, err := t.db.Exec(constant.TRANSACTION_INSERT, bodyRequest.Id, bodyRequest.CustomerId, bodyRequest.MerchantId, bodyRequest.BankId, bodyRequest.Amount, bodyRequest.TransactionAt); err != nil {
		return err
	}

	return nil
}

func (t *transactionRepository) List() ([]model.Transaction, error) {
	rows, err := t.db.Query(constant.TRANSACTION_LIST)
	if err != nil {
		return nil, err
	}

	transactions, err := t.Scan(rows)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (t *transactionRepository) History(customerId string) ([]model.Transaction, error) {
	rows, err := t.db.Query(constant.TRANSACTION_HISTORY, customerId)
	if err != nil {
		return nil, err
	}

	transactions, err := t.Scan(rows)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (t *transactionRepository) Scan(rows *sql.Rows) ([]model.Transaction, error) {
	var transactions []model.Transaction
	for rows.Next() {
		var transaction model.Transaction
		if err := rows.Scan(&transaction.Id, &transaction.CustomerId, &transaction.MerchantId, &transaction.BankId, &transaction.Amount, &transaction.TransactionAt); err != nil {
			return nil, err
		}

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func NewTransactionRepository(db *sql.DB) TransactionRepository {
	return &transactionRepository{db}
}
