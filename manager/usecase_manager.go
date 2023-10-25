package manager

type UseCaseManager interface {
}

type useCaseManager struct {
	repoManager RepoManager
}

func NewUseCaseManager(repo RepoManager) UseCaseManager {
	return &useCaseManager{
		repoManager: repo,
	}
}
