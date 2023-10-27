package model

import (
	"time"
)

type Transaction struct {
	Id            string `json:"id" binding:"required"`
	CustomerId    string `json:"customer_id" binding:"required"`
	MerchantId    string `json:"merchant_id" binding:"required"`
	BankId        string `json:"bank_id" binding:"required"`
	Amount        uint64 `json:"amount" binding:"required"`
	TransactionAt time.Time
}

type TransactionDTO struct {
	Id            string    `json:"id"`
	Customer      Customer  `json:"customer"`
	Merchant      Merchant  `json:"merchant"`
	Bank          Bank      `json:"bank"`
	Amount        uint64    `json:"amount"`
	TransactionAt time.Time `json:"transaction_at"`
}
