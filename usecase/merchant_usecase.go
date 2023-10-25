package usecase

import (
	"fmt"
	"go-transaction/model"
	"go-transaction/repository"
)

type MerchantUsecase interface {
	RegisterNewMerchant(bodyRequest model.Merchant) error
	ShowListMerchant() ([]model.Merchant, error)
	GetDetailmerchant(id string) (model.Merchant, error)
	EditExistedMerchant(bodyRequest model.Merchant) error
}

type merchantUsecase struct {
	repo repository.MerchantRepository
}

func (m *merchantUsecase) RegisterNewMerchant(bodyRequest model.Merchant) error {
	err := m.repo.Create(bodyRequest)

	if err != nil {
		return fmt.Errorf("failed to add merchant : %s", err.Error())
	}
	return nil
}

func (m *merchantUsecase) ShowListMerchant() ([]model.Merchant, error) {
	return m.repo.List()
}

func (m *merchantUsecase) GetDetailmerchant(id string) (model.Merchant, error) {
	return m.repo.Detail(id)
}

func (m *merchantUsecase) EditExistedMerchant(bodyRequest model.Merchant) error {
	if _, err := m.GetDetailmerchant(bodyRequest.Id); err != nil {
		return fmt.Errorf("merchant is not exist")
	}

	if err := m.repo.Update(bodyRequest); err != nil {
		return fmt.Errorf("failed to update merchant : %s", err.Error())
	}

	return nil
}

func NewMerchantUsecase(repo repository.MerchantRepository) MerchantUsecase {
	return &merchantUsecase{repo}
}
