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

func (repository *Repository) GetStatus() (models.Status, error) {
	var result models.Status
	row := repository.db.QueryRow(
		`SELECT * FROM
		(SELECT COUNT(*) FROM users) as u,
 		(SELECT COUNT(*) FROM forum) as f,
		(SELECT COUNT(*) FROM thread) as t,
		(SELECT COUNT(*) FROM post) as p;`)

	err := row.Scan(
		&result.User,
		&result.Forum,
		&result.Thread,
		&result.Post)
	if err != nil {
		return models.Status{}, err
	}

	return result, nil
}
func (repository *Repository) Clear() error {
	_, err := repository.db.Exec(`TRUNCATE users, forum, thread, post, vote, users_forum;`)
	if err != nil {
		return err
	}

	return nil
}