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

func (uc *UseCase) GetUserProfile(nickname string) (models.User, *models.CustomError) {
	user, err := uc.Repository.GetUser(nickname)
	if err == pgx.ErrNoRows {
		return models.User{}, &models.CustomError{Message: models.NoUser}
	}
	return user, nil
}

func (uc *UseCase) UpdateUserProfile(user models.User) (models.User, *models.CustomError) {
	userNew, err := uc.Repository.UpdateUser(user)
	if err != nil {
		if pgErr, ok := err.(pgx.PgError); ok && pgErr.Code == "23505" {
			return models.User{}, &models.CustomError{Message: models.ConflictData}
		}
		if err == pgx.ErrNoRows {
			return models.User{}, &models.CustomError{Message: models.NoUser}
		}

		return models.User{}, &models.CustomError{Message: err.Error()}
	}
	return userNew, nil
}
