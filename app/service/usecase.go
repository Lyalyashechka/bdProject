package service

import "github.com/Lyalyashechka/bdProject/app/models"

type UseCase interface {
	GetStatus() (models.Status, error)
	Clear() error
}