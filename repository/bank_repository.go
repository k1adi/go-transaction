package repository

import (
	"database/sql"
	"go-transaction/model"
	"go-transaction/utils/constant"
)

type BankRepository interface {
	BaseRepository[model.Bank]
}

type bankRepository struct {
	db *sql.DB
}

func (b *bankRepository) Create(bodyRequest model.Bank) error {
	if _, err := b.db.Exec(constant.BANK_INSERT, bodyRequest.Id, bodyRequest.Name); err != nil {
		return err
	}

	return nil
}

func (b *bankRepository) List() ([]model.Bank, error) {
	rows, err := b.db.Query(constant.BANK_LIST)
	if err != nil {
		return nil, err
	}

	banks, err := b.Scan(rows)
	if err != nil {
		return nil, err
	}
	return banks, nil
}

func (b *bankRepository) Scan(rows *sql.Rows) ([]model.Bank, error) {
	var banks []model.Bank
	for rows.Next() {
		var bank model.Bank
		if err := rows.Scan(&bank.Id, &bank.Name); err != nil {
			return nil, err
		}

		banks = append(banks, bank)
	}

	return banks, nil
}

func (b *bankRepository) Detail(id string) (model.Bank, error) {
	var bank model.Bank

	if err := b.db.QueryRow(constant.BANK_DETAIL, id).Scan(&bank.Id, &bank.Name); err != nil {
		return model.Bank{}, err
	}

	return bank, nil
}

func (b *bankRepository) Update(bodyRequest model.Bank) error {
	if _, err := b.db.Exec(constant.BANK_UPDATE, bodyRequest.Name, bodyRequest.Id); err != nil {
		return err
	}

	return nil
}

func NewBankRepository(db *sql.DB) BankRepository {
	return &bankRepository{db}
}
