package manager

import (
	"portfolio-api/repository"
)

type RepoManager interface {
	VisitsRepository() repository.VisitsRepository
}

type repoManager struct {
	infraManager InfraManager
}

func (r *repoManager) VisitsRepository() repository.VisitsRepository {
	return repository.NewVisitsRepository(r.infraManager.GetDB())
}

func NewRepoManager(infraManager InfraManager) RepoManager {
	return &repoManager{infraManager: infraManager}
}