package manager

type RepoManager interface {
}

type repoManager struct {
	infra InfraManager
}

func NewRepoManager(infraParam InfraManager) RepoManager {
	return &repoManager{
		infra: infraParam,
	}
}
