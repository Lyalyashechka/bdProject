package service

import "github.com/Lyalyashechka/bdProject/app/models"

type Repository interface {
	GetStatus() (models.Status, error)
	Clear() error
}