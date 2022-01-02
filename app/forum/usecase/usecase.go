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

func (uc *UseCase) CreateThread (thread models.Thread) (models.Thread, *models.CustomError) {
	thread, err := uc.Repository.AddThread(thread)
	if err != nil {
		if pgErr, ok := err.(pgx.PgError); ok && pgErr.Code == models.PgxNoFoundFieldErrorCode {
			return models.Thread{}, &models.CustomError{Message: models.NoUser}
		}
		if pgErr, ok := err.(pgx.PgError); ok && pgErr.Code == models.PgxUniqErrorCode {
			return models.Thread{}, &models.CustomError{Message: models.ConflictData}
		}
		return models.Thread{}, &models.CustomError{Message: err.Error()}
	}

	return thread, nil
}

func (uc *UseCase) GetUsersForum (slug string) ([]models.User, *models.CustomError) {
	users, err := uc.Repository.GetUsersForum(slug)
	if users == nil {
		return nil, &models.CustomError{Message: models.NoSlug}
	}
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &models.CustomError{Message: models.NoSlug}
		}
		return nil, &models.CustomError{Message: err.Error()}
	}

	return users, nil
}

func (uc *UseCase) GetForumThreads (slug string) ([]models.Thread, *models.CustomError) {
	threads, err := uc.Repository.GetForumThreads(slug)
	if threads == nil {
		return nil, &models.CustomError{Message: models.NoSlug}
	}
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &models.CustomError{Message: models.NoSlug}
		}
		return nil, &models.CustomError{Message: err.Error()}
	}

	return threads, nil
}