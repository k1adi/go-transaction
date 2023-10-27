package usecase

import (
	"fmt"
	"go-transaction/model"
	"go-transaction/repository"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type CustomerUsecase interface {
	RegisterNewCustomer(bodyRequest model.Customer) error
	ShowListCustomers() ([]model.Customer, error)
	GetDetailCustomer(id string) (model.Customer, error)
	FindUsernameAndPassword(username, password string) (model.Auth, error)
}

type customerUsecase struct {
	repo repository.CustomerRepository
}

func (c *customerUsecase) RegisterNewCustomer(bodyRequest model.Customer) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(bodyRequest.Password), 14)
	if err != nil {
		return err
	}

	bodyRequest.Password = string(bytes)
	bodyRequest.CreatedAt = time.Now()
	err = c.repo.Create(bodyRequest)

	if err != nil {
		return fmt.Errorf("failed to add customer : %s", err.Error())
	}
	return nil
}

func (c *customerUsecase) ShowListCustomers() ([]model.Customer, error) {
	return c.repo.List()
}

func (c *customerUsecase) FindUsernameAndPassword(username, password string) (model.Auth, error) {
	customer, err := c.repo.FindUsername(username)
	if err != nil {
		return model.Auth{}, fmt.Errorf("error find customer username : %v", err.Error())
	}

	if err = bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(password)); err != nil {
		return model.Auth{}, err
	}

	return customer, nil
}

func (c *customerUsecase) GetDetailCustomer(id string) (model.Customer, error) {
	return c.repo.Detail(id)
}

func NewCustomerUsecase(repo repository.CustomerRepository) CustomerUsecase {
	return &customerUsecase{repo}
}
