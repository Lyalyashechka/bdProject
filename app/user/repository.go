package user

import "github.com/Lyalyashechka/bdProject/app/models"

type Repository interface {
	AddUser(user models.User) (models.User, error)
	GetUser(nickname string) (models.User, error)
	UpdateUser(nickname string, update models.UserUpdate) (models.User, error)
	GetUsersByNicknameOrEmail(nickname string, email string) ([]models.User, error)
}
