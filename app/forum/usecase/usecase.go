package usecase

import (
	"database/sql"
	"github.com/Lyalyashechka/bdProject/app/forum"
	"github.com/Lyalyashechka/bdProject/app/models"
	"github.com/jackc/pgx"
)

type UseCase struct {
	Repository forum.Repository
}

func NewUseCase(repository forum.Repository) *UseCase {
	return &UseCase{Repository: repository}
}

func (uc *UseCase) CreateForum(forum models.Forum) (models.Forum, *models.CustomError) {
	forum, err := uc.Repository.AddForum(forum)
	if err != nil {
		if pgErr, ok := err.(pgx.PgError); ok && pgErr.Code == models.PgxNoFoundFieldErrorCode {
			return models.Forum{}, &models.CustomError{Message: models.NoUser}
		}
		if pgErr, ok := err.(pgx.PgError); ok && pgErr.Code == models.PgxUniqErrorCode {
			return models.Forum{}, &models.CustomError{Message: models.ConflictData}
		}
		return models.Forum{}, &models.CustomError{Message: err.Error()}
	}

	return forum, nil
}

func (uc *UseCase) GetDetailsForum(slug string) (models.Forum, *models.CustomError) {
	forum, err := uc.Repository.GetDetailsForum(slug)
	if err == sql.ErrNoRows {
		return models.Forum{}, &models.CustomError{Message: models.NoSlug}
	}
	return forum, nil
}
