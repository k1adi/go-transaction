package manager

import "go-transaction/repository"

type RepoManager interface {
	BankRepository() repository.BankRepository
	MerchantRepository() repository.MerchantRepository
}

type repoManager struct {
	infra InfraManager
}

func (r *repoManager) BankRepository() repository.BankRepository {
	return repository.NewBankRepository(r.infra.Connection())
}

func (r *repoManager) MerchantRepository() repository.MerchantRepository {
	return repository.NewMerchantRepository(r.infra.Connection())
}

func NewRepoManager(infraParam InfraManager) RepoManager {
	return &repoManager{
		infra: infraParam,
	}
}
