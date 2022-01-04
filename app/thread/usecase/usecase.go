package usecase

import (
	"database/sql"
	"github.com/Lyalyashechka/bdProject/app/models"
	"github.com/Lyalyashechka/bdProject/app/thread"
	"github.com/jackc/pgx"
	"strconv"
)

type UseCase struct {
	Repository thread.Repository
}

func NewUseCase(repository thread.Repository) *UseCase {
	return &UseCase{Repository: repository}
}

func (uc *UseCase) CreatePosts(slugOrId string, post []models.Post) ([]models.Post, *models.CustomError) {
	var thread models.Thread
	id, err := strconv.Atoi(slugOrId)
	if err != nil {
		thread, err = uc.Repository.GetThreadBySlug(slugOrId)
		if err != nil {
			if pgErr, ok := err.(pgx.PgError); ok && pgErr.Code == "23505" {
				return nil, &models.CustomError{Message: models.ConflictData}
			}

			if err == sql.ErrNoRows {
				return nil, &models.CustomError{Message: models.NoUser}
			}

			return nil, &models.CustomError{Message: err.Error()}
		}

	} else {
		thread, err = uc.Repository.GetThreadById(id)
		if err != nil {
			if pgErr, ok := err.(pgx.PgError); ok && pgErr.Code == "23505" {
				return nil, &models.CustomError{Message: models.ConflictData}
			}

			if err == sql.ErrNoRows {
				return nil, &models.CustomError{Message: models.NoUser}
			}

			return nil, &models.CustomError{Message: err.Error()}
		}
	}
	posts, err := uc.Repository.CreatePosts(int(thread.Id), thread.Forum, post)
	if err != nil {
		if pgErr, ok := err.(pgx.PgError); ok && pgErr.Code == "23505" {
			return nil, &models.CustomError{Message: models.ConflictData}
		}
		if err == sql.ErrNoRows {
			return nil, &models.CustomError{Message: models.NoUser}
		}
		return nil, &models.CustomError{Message: err.Error()}
	}

	return posts, nil
}

func (uc *UseCase) CreateVote(slugOrId string, vote models.Vote) (models.Thread, *models.CustomError) {
	var thread models.Thread
	err := uc.Repository.CreateVoteBySlugOrId(slugOrId, vote)
	if err != nil {
		if pgErr, ok := err.(pgx.PgError); ok && pgErr.Code == "23503" {
			return models.Thread{}, &models.CustomError{Message: models.NoUser}
		}
		if pgErr, ok := err.(pgx.PgError); ok && pgErr.Code == "23505" {
			err = uc.Repository.UpdateVoteBySlugOrId(slugOrId, vote)
			if err != nil {
				return models.Thread{}, &models.CustomError{Message: err.Error()}
			}
			thread, err = uc.Repository.GetThreadBySlugOrId(slugOrId)
			if err != nil {
				return models.Thread{}, &models.CustomError{Message: err.Error()}
			}

			return thread, nil
		}
		return models.Thread{}, &models.CustomError{Message: err.Error()}
	}

	thread, err = uc.Repository.GetThreadBySlugOrId(slugOrId)
	if err != nil {
		return models.Thread{}, &models.CustomError{Message: err.Error()}
	}

	return thread, nil
}

func (uc *UseCase) GetThreadDetails(slugOrId string) (models.Thread, *models.CustomError) {
	thread, err := uc.Repository.GetThreadBySlugOrId(slugOrId)
	if err != nil {
		return models.Thread{}, &models.CustomError{Message: err.Error()}
	}
	return thread, nil
}
