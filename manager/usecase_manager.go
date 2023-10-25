package manager

import "go-transaction/usecase"

type UseCaseManager interface {
	BankUsecase() usecase.BankUsecase
	MerchantUsecase() usecase.MerchantUsecase
}

type useCaseManager struct {
	repoManager RepoManager
}

func (u *useCaseManager) BankUsecase() usecase.BankUsecase {
	return usecase.NewBankUsecase(u.repoManager.BankRepository())
}

func (u *useCaseManager) MerchantUsecase() usecase.MerchantUsecase {
	return usecase.NewMerchantUsecase(u.repoManager.MerchantRepository())
}

func NewUseCaseManager(repo RepoManager) UseCaseManager {
	return &useCaseManager{
		repoManager: repo,
	}
}
