package usecase

import (
	"fmt"
	"go-transaction/model"
	"go-transaction/repository"
)

type BankUsecase interface {
	RegisterNewBank(bodyRequest model.Bank) error
	ShowListBank() ([]model.Bank, error)
	GetDetailBank(id string) (model.Bank, error)
	EditExistedBank(bodyRequest model.Bank) error
}

type bankUsecase struct {
	repo repository.BankRepository
}

func (b *bankUsecase) RegisterNewBank(bodyRequest model.Bank) error {
	err := b.repo.Create(bodyRequest)

	if err != nil {
		return fmt.Errorf("failed to add bank : %s", err.Error())
	}
	return nil
}

func (b *bankUsecase) ShowListBank() ([]model.Bank, error) {
	return b.repo.List()
}

func (b *bankUsecase) GetDetailBank(id string) (model.Bank, error) {
	return b.repo.Detail(id)
}

func (b *bankUsecase) EditExistedBank(bodyRequest model.Bank) error {
	if _, err := b.GetDetailBank(bodyRequest.Id); err != nil {
		return fmt.Errorf("bank id is not exist")
	}

	if err := b.repo.Update(bodyRequest); err != nil {
		return fmt.Errorf("failed to update bank : %s", err.Error())
	}

	return nil
}

func NewBankUsecase(repo repository.BankRepository) BankUsecase {
	return &bankUsecase{repo}
}
