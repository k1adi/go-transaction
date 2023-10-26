package manager

import "go-transaction/usecase"

type UseCaseManager interface {
	BankUsecase() usecase.BankUsecase
	MerchantUsecase() usecase.MerchantUsecase
	CustomerUsecase() usecase.CustomerUsecase
	AdminUsecase() usecase.AdminUsecase
	AuthUsecase() usecase.AuthUsecase
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

func (u *useCaseManager) CustomerUsecase() usecase.CustomerUsecase {
	return usecase.NewCustomerUsecase(u.repoManager.CustomerRepository())
}

func (u *useCaseManager) AdminUsecase() usecase.AdminUsecase {
	return usecase.NewAdminUsecase(u.repoManager.AdminRepository())
}

func (u *useCaseManager) AuthUsecase() usecase.AuthUsecase {
	return usecase.NewAuthUsecase(u.CustomerUsecase(), u.AdminUsecase())
}

func NewUseCaseManager(repo RepoManager) UseCaseManager {
	return &useCaseManager{
		repoManager: repo,
	}
}
