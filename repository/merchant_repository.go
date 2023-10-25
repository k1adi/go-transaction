package repository

import (
	"database/sql"
	"go-transaction/model"
	"go-transaction/utils/constant"
)

type MerchantRepository interface {
	BaseRepository[model.Merchant]
}

type merchantRepository struct {
	db *sql.DB
}

func (m *merchantRepository) Create(bodyRequest model.Merchant) error {
	if _, err := m.db.Exec(constant.MERCHANT_INSERT, bodyRequest.Id, bodyRequest.Name, bodyRequest.Address); err != nil {
		return err
	}

	return nil
}

func (m *merchantRepository) List() ([]model.Merchant, error) {
	rows, err := m.db.Query(constant.MERCHANT_LIST)
	if err != nil {
		return nil, err
	}

	merchants, err := m.Scan(rows)
	if err != nil {
		return nil, err
	}
	return merchants, nil
}

func (m *merchantRepository) Scan(rows *sql.Rows) ([]model.Merchant, error) {
	var merchants []model.Merchant
	for rows.Next() {
		var merchant model.Merchant
		if err := rows.Scan(&merchant.Id, &merchant.Name, &merchant.Address); err != nil {
			return nil, err
		}

		merchants = append(merchants, merchant)
	}

	return merchants, nil
}

func (m *merchantRepository) Detail(id string) (model.Merchant, error) {
	var merchant model.Merchant

	if err := m.db.QueryRow(constant.MERCHANT_DETAIL, id).Scan(&merchant.Id, &merchant.Name, &merchant.Address); err != nil {
		return model.Merchant{}, err
	}

	return merchant, nil
}

func (m *merchantRepository) Update(bodyRequest model.Merchant) error {
	if _, err := m.db.Exec(constant.MERCHANT_UPDATE, bodyRequest.Name, bodyRequest.Address, bodyRequest.Id); err != nil {
		return err
	}

	return nil
}

func NewMerchantRepository(db *sql.DB) MerchantRepository {
	return &merchantRepository{db}
}
