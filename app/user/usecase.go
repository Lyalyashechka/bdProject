package user

import "github.com/Lyalyashechka/bdProject/app/models"

type UseCase interface {
	CreateUser(user models.User) ([]models.User, error)
	GetUserProfile(nickname string) (models.User, *models.CustomError)
	UpdateUserProfile(user models.User) (models.User, *models.CustomError)
}
