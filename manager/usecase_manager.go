package manager

import "go-transaction/usecase"

type UseCaseManager interface {
	BankUsecase() usecase.BankUsecase
}

type useCaseManager struct {
	repoManager RepoManager
}

func (u *useCaseManager) BankUsecase() usecase.BankUsecase {
	return usecase.NewBankUsecase(u.repoManager.BankRepository())
}

func NewUseCaseManager(repo RepoManager) UseCaseManager {
	return &useCaseManager{
		repoManager: repo,
	}
}
