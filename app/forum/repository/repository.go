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

func (repository *Repository) AddThread(thread models.Thread) (models.Thread, error){
	_, err := repository.db.Exec("INSERT INTO thread (title, author, forum, message, slug) " +
		"VALUES ($1, $2, $3, $4, $5)",
		thread.Title, thread.Author, thread.Forum, thread.Message, thread.Slug)
	if err != nil {
		return models.Thread{}, err
	}

	return thread, nil
}

func (repository *Repository) GetUsersForum(slug string)([]models.User, error) {
	rows, err := repository.db.Query("SELECT u.nickname, fullname, about, email "+
		"FROM users_forum as u inner join users on u.nickname = users.nickname where u.slug = $1", slug)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.Nickname, &user.FullName, &user.About, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return users, nil
}

func (repository *Repository) GetForumThreads (slug string) ([]models.Thread, error) {
	rows, err := repository.db.Query("select id, title, author, forum, message, votes, slug, created " +
		"from thread where forum = $1", slug)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var threads []models.Thread
	for rows.Next() {
		var thread models.Thread
		err = rows.Scan(&thread.Id, &thread.Title, &thread.Author, &thread.Forum, &thread.Message,
			&thread.Votes, &thread.Slug, &thread.Created)
		if err != nil {
			return nil, err
		}
		threads = append(threads, thread)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return threads, nil
}
