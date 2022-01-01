package usecase

import (
	"github.com/Lyalyashechka/bdProject/app/models"
	"github.com/Lyalyashechka/bdProject/app/user"
	"github.com/jackc/pgx"
)

type UseCase struct {
	Repository user.Repository
}

func NewUseCase(repository user.Repository) *UseCase {
	return &UseCase{Repository: repository}
}

func (uc *UseCase) CreateUser(user models.User) ([]models.User, error) {
	var resultArray []models.User
	result, err := uc.Repository.AddUser(user)
	if err != nil {
		if pgErr, ok := err.(pgx.PgError); ok && pgErr.Code == "23505" {
			result, err1 := uc.Repository.GetUsersByNicknameOrEmail(user.Nickname, user.Email)
			if err1 != nil {
				return nil, err1
			}
			return result, err
		}
	}

	resultArray = append(resultArray, result)
	return resultArray, err
}

func (uc *UseCase) GetUserProfile(nickname string) (models.User, error) {
	return models.User{}, nil
}

func (uc *UseCase) UpdateUserProfile(nickname string, update models.UserUpdate) (models.User, error) {
	return models.User{}, nil
}
