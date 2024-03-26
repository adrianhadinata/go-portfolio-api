package usecase

import (
	"portfolio-api/entity"
	_ "portfolio-api/entity/dto"
	"portfolio-api/repository"
)

type VisitsUsecase interface {
	Insert(visitor entity.Visitor) (entity.Visitor, error)
}

type visitsUsecase struct {
	visitsRepository repository.VisitsRepository
}

func (v *visitsUsecase) Insert(visitor entity.Visitor) (entity.Visitor, error) {
	visitor, err := v.visitsRepository.Insert(visitor)
	if err != nil {
		return entity.Visitor{}, err
	}

	return visitor, nil
}

func NewVisitsUsecase(visitsRepository repository.VisitsRepository) VisitsUsecase {
	return &visitsUsecase{visitsRepository: visitsRepository}
}