package manager

import "go-transaction/repository"

type RepoManager interface {
	BankRepository() repository.BankRepository
}

type repoManager struct {
	infra InfraManager
}

func (r *repoManager) BankRepository() repository.BankRepository {
	return repository.NewBankRepository(r.infra.Connection())
}

func NewRepoManager(infraParam InfraManager) RepoManager {
	return &repoManager{
		infra: infraParam,
	}
}
