package manager

import "go-transaction/repository"

type RepoManager interface {
	BankRepository() repository.BankRepository
	MerchantRepository() repository.MerchantRepository
	AdminRepository() repository.AdminRepository
	CustomerRepository() repository.CustomerRepository
	TransactionRepository() repository.TransactionRepository
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

func (r *repoManager) AdminRepository() repository.AdminRepository {
	return repository.NewAdminRepository(r.infra.Connection())
}

func (r *repoManager) CustomerRepository() repository.CustomerRepository {
	return repository.NewCustomerRepository(r.infra.Connection())
}

func (r *repoManager) TransactionRepository() repository.TransactionRepository {
	return repository.NewTransactionRepository(r.infra.Connection())
}

func NewRepoManager(infraParam InfraManager) RepoManager {
	return &repoManager{
		infra: infraParam,
	}
}
