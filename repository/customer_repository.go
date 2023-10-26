package repository

import (
	"database/sql"
	"go-transaction/model"
	"go-transaction/utils/constant"
)

type CustomerRepository interface {
	BaseRepository[model.Customer]
	FindUsername(username string) (model.Auth, error)
}

type customerRepository struct {
	db *sql.DB
}

func (c *customerRepository) Create(bodyRequest model.Customer) error {
	if _, err := c.db.Exec(constant.CUSTOMER_INSERT, bodyRequest.Id, bodyRequest.FullName, bodyRequest.UserName, bodyRequest.Password, bodyRequest.PhoneNumber, bodyRequest.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (c *customerRepository) List() ([]model.Customer, error) {
	rows, err := c.db.Query(constant.CUSTOMER_LIST)
	if err != nil {
		return nil, err
	}

	customers, err := c.Scan(rows)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (c *customerRepository) Scan(rows *sql.Rows) ([]model.Customer, error) {
	var customers []model.Customer
	for rows.Next() {
		var customer model.Customer
		if err := rows.Scan(&customer.Id, &customer.FullName, &customer.UserName, &customer.PhoneNumber, &customer.CreatedAt); err != nil {
			return nil, err
		}

		customers = append(customers, customer)
	}

	return customers, nil
}

func (c *customerRepository) Detail(id string) (model.Customer, error) {
	var customer model.Customer

	if err := c.db.QueryRow(constant.CUSTOMER_DETAIL, id).Scan(&customer.Id, &customer.FullName, &customer.UserName, &customer.Password, &customer.CreatedAt); err != nil {
		return model.Customer{}, err
	}

	return customer, nil
}

func (c *customerRepository) Update(bodyRequest model.Customer) error {
	if _, err := c.db.Exec(constant.CUSTOMER_UPDATE, bodyRequest.FullName, bodyRequest.UserName, bodyRequest.PhoneNumber, bodyRequest.Id); err != nil {
		return err
	}

	return nil
}

func (c *customerRepository) FindUsername(username string) (model.Auth, error) {
	var customer model.Auth
	if err := c.db.QueryRow(constant.CUSTOMER_VALIDATION, username).Scan(&customer.Id, &customer.Username, &customer.Password); err != nil {
		return model.Auth{}, err
	}

	return customer, nil
}

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	return &customerRepository{db}
}
