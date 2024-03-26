package manager

import (
	"portfolio-api/usecase"
)

type UsecaseManager interface {
	VisitsUsecase() usecase.VisitsUsecase
}

type usecaseManager struct {
	repository RepoManager
}

func (u *usecaseManager) VisitsUsecase() usecase.VisitsUsecase {
	return usecase.NewVisitsUsecase(u.repository.VisitsRepository())
}

func NewUseCaseManager(repository RepoManager) UsecaseManager {
	return &usecaseManager{repository: repository}
}