package usecase

import (
	"github.com/Lyalyashechka/bdProject/app/models"
	"github.com/Lyalyashechka/bdProject/app/service"
)

type UseCase struct {
	Repository service.Repository
}

func NewUseCase(repository service.Repository) *UseCase {
	return &UseCase{Repository: repository}
}

func (uc *UseCase) GetStatus() (models.Status, error)  {
	return uc.Repository.GetStatus()
}

func (uc *UseCase) Clear() error {
	return uc.Repository.Clear()
}