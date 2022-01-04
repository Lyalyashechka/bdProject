package usecase

import (
	"database/sql"
	"github.com/Lyalyashechka/bdProject/app/forum"
	"github.com/Lyalyashechka/bdProject/app/models"
	thread "github.com/Lyalyashechka/bdProject/app/thread"
	"github.com/Lyalyashechka/bdProject/app/tools"
	"github.com/google/uuid"
	"github.com/jackc/pgx"
)

type UseCase struct {
	RepositoryForum  forum.Repository
	RepositoryThread thread.Repository
}

func NewUseCase(repositoryForum forum.Repository, repository thread.Repository) *UseCase {
	return &UseCase{RepositoryForum: repositoryForum, RepositoryThread: repository}
}

func (uc *UseCase) CreateForum(forumGet models.Forum) (models.Forum, *models.CustomError) {
	forum, err := uc.RepositoryForum.AddForum(forumGet)
	if err != nil {
		if pgErr, ok := err.(pgx.PgError); ok && pgErr.Code == models.PgxNoFoundFieldErrorCode {
			return models.Forum{}, &models.CustomError{Message: models.NoUser}
		}
		if pgErr, ok := err.(pgx.PgError); ok && pgErr.Code == models.PgxUniqErrorCode {
			forum, err = uc.RepositoryForum.GetForumBySlug(forumGet.Slug)
			if err != nil {
				return models.Forum{}, &models.CustomError{Message: err.Error()}
			}
			return forum, &models.CustomError{Message: models.ConflictData}
		}
		return models.Forum{}, &models.CustomError{Message: err.Error()}
	}

	return forum, nil
}

func (uc *UseCase) GetDetailsForum(slug string) (models.Forum, *models.CustomError) {
	forum, err := uc.RepositoryForum.GetDetailsForum(slug)
	if err == sql.ErrNoRows {
		return models.Forum{}, &models.CustomError{Message: models.NoSlug}
	}
	return forum, nil
}

func (uc *UseCase) CreateThread(threadGet models.Thread) (models.Thread, *models.CustomError) {
	var randomSlug bool
	if threadGet.Slug == "" {
		randomSlug = true
		threadGet.Slug = uuid.NewString()
	}
	thread, err := uc.RepositoryForum.AddThread(threadGet)
	if err != nil {
		if pgErr, ok := err.(pgx.PgError); ok && pgErr.Code == models.PgxNoFoundFieldErrorCode {
			return models.Thread{}, &models.CustomError{Message: models.NoUser}
		}
		if pgErr, ok := err.(pgx.PgError); ok && pgErr.Code == models.PgxUniqErrorCode {
			thread, err = uc.RepositoryThread.GetThreadBySlug(threadGet.Slug)
			if err != nil {
				return models.Thread{}, &models.CustomError{Message: err.Error()}
			}
			return thread, &models.CustomError{Message: models.ConflictData}
		}
		return models.Thread{}, &models.CustomError{Message: err.Error()}
	}

	if randomSlug == true {
		thread.Slug = ""
	}
	return thread, nil
}

func (uc *UseCase) GetUsersForum(slug string) ([]models.User, *models.CustomError) {
	users, err := uc.RepositoryForum.GetUsersForum(slug)
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

func (uc *UseCase) GetForumThreads(slug string, filter tools.Filter) ([]models.Thread, *models.CustomError) {
	threads, err := uc.RepositoryForum.GetForumThreads(slug, filter)
	if threads == nil {
		_, err := uc.RepositoryForum.GetForumBySlug(slug)
		if err != nil {
			return nil, &models.CustomError{Message: err.Error()}
		}
		return []models.Thread{}, nil
	}
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &models.CustomError{Message: models.NoSlug}
		}
		return nil, &models.CustomError{Message: err.Error()}
	}

	return threads, nil
}
