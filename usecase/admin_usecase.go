package usecase

import (
	"fmt"
	"go-transaction/model"
	"go-transaction/repository"

	"golang.org/x/crypto/bcrypt"
)

type AdminUsecase interface {
	RegisterNewAdmin(bodyRequest model.Auth) error
	ShowListAdmins() ([]model.Auth, error)
	FindUsernameAndPassword(username, password string) (model.Auth, error)
}

type adminUsecase struct {
	repo repository.AdminRepository
}

func (a *adminUsecase) RegisterNewAdmin(bodyRequest model.Auth) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(bodyRequest.Password), 14)
	if err != nil {
		return err
	}

	bodyRequest.Password = string(bytes)
	err = a.repo.Create(bodyRequest)

	if err != nil {
		return fmt.Errorf("failed to add admin : %s", err.Error())
	}
	return nil
}

func (a *adminUsecase) ShowListAdmins() ([]model.Auth, error) {
	return a.repo.List()
}

func (a *adminUsecase) FindUsernameAndPassword(username, password string) (model.Auth, error) {
	customer, err := a.repo.FindUsername(username)
	if err != nil {
		return model.Auth{}, fmt.Errorf("error find customer username : %v", err.Error())
	}

	if err = bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(password)); err != nil {
		return model.Auth{}, err
	}

	return customer, nil
}

func NewAdminUsecase(repo repository.AdminRepository) AdminUsecase {
	return &adminUsecase{repo}
}
