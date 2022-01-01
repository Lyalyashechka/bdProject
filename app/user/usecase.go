package user

import "github.com/Lyalyashechka/bdProject/app/models"

type UseCase interface {
	CreateUser(user models.User) ([]models.User, error)
	GetUserProfile(nickname string) (models.User, error)
	UpdateUserProfile(nickname string, update models.UserUpdate) (models.User, error)
}
