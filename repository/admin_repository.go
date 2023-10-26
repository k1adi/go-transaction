package repository

import (
	"database/sql"
	"go-transaction/model"
	"go-transaction/utils/constant"
)

type AdminRepository interface {
	Create(bodyRequest model.Auth) error
	List() ([]model.Auth, error)
	Scan(rows *sql.Rows) ([]model.Auth, error)
	FindUsername(username string) (model.Auth, error)
}

type adminRepository struct {
	db *sql.DB
}

func (a *adminRepository) Create(bodyRequest model.Auth) error {
	if _, err := a.db.Exec(constant.ADMIN_INSERT, bodyRequest.Id, bodyRequest.Username, bodyRequest.Password); err != nil {
		return err
	}

	return nil
}

func (a *adminRepository) List() ([]model.Auth, error) {
	rows, err := a.db.Query(constant.ADMIN_LIST)
	if err != nil {
		return nil, err
	}

	admins, err := a.Scan(rows)
	if err != nil {
		return nil, err
	}
	return admins, nil
}

func (a *adminRepository) Scan(rows *sql.Rows) ([]model.Auth, error) {
	var admins []model.Auth
	for rows.Next() {
		var admin model.Auth
		if err := rows.Scan(&admin.Id, &admin.Username); err != nil {
			return nil, err
		}

		admins = append(admins, admin)
	}

	return admins, nil
}

func (a *adminRepository) FindUsername(username string) (model.Auth, error) {
	var admin model.Auth
	if err := a.db.QueryRow(constant.ADMIN_VALIDATION, username).Scan(&admin.Id, &admin.Username, &admin.Password); err != nil {
		return model.Auth{}, err
	}

	return admin, nil
}

func NewAdminRepository(db *sql.DB) AdminRepository {
	return &adminRepository{db}
}
