package repository

import (
	"database/sql"
	"github.com/Lyalyashechka/bdProject/app/models"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (repository *Repository) AddForum(forum models.Forum) (models.Forum, error) {
	_, err := repository.db.Exec("INSERT INTO Forum (title, \"user\", slug) VALUES ($1, $2, $3)",
		forum.Title, forum.User, forum.Slug)
	if err != nil {
		return models.Forum{}, err
	}

	return forum, nil
}

func (repository *Repository) GetDetailsForum(slug string) (models.Forum, error) {
	var result models.Forum
	row := repository.db.QueryRow("SELECT title, \"user\", slug, posts, threads "+
		"FROM Forum WHERE slug=$1", slug)

	err := row.Scan(&result.Title, &result.User, &result.Slug, &result.Posts, &result.Threads)
	if err != nil {
		return models.Forum{}, err
	}
	return result, nil
}
