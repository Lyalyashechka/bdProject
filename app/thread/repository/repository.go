package repository

import (
	"database/sql"
	"fmt"
	"github.com/Lyalyashechka/bdProject/app/models"
	"strings"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (repository *Repository)CreatePosts (threadId int,threadForum string, posts []models.Post)([]models.Post, error) {
	query := `INSERT INTO post(parent, author, message, thread, forum) VALUES `
	var values []interface{}

	for i, post := range posts {
		value := fmt.Sprintf("($%d, $%d, $%d, $%d, $%d),",
			i*5+1,i*5+2,i*5+3,i*5+4,i*5+5)
		query += value
		values = append(values, post.Parent, post.Author, post.Message, threadId, threadForum)
	}
	query = strings.TrimSuffix(query, ",")
	query += ` RETURNING id, parent, author, message, isEdited, forum, thread, created;`

	rows, err := repository.db.Query(query, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.Id, &post.Parent, &post.Author, &post.Message,
			&post.IsEdited, &post.Forum, &post.Thread, &post.Created)
		if err != nil {
			return nil, err
		}
		result = append(result, post)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return result, nil
}

func (repository *Repository) GetThreadBySlug(slug string)(models.Thread, error){
	var result models.Thread
	row := repository.db.QueryRow("SELECT id, title, author, forum, message, votes, slug, created "+
		"FROM thread WHERE slug=$1", slug)

	err := row.Scan(&result.Id, &result.Title, &result.Author, &result.Forum, &result.Message, &result.Votes,
		&result.Slug, &result.Created)
	if err != nil {
		return models.Thread{}, err
	}

	return result, nil
}

func (repository *Repository) GetThreadById(id int)(models.Thread, error) {
	var result models.Thread
	row := repository.db.QueryRow("SELECT id, title, author, forum, message, votes, slug, created "+
		"FROM thread WHERE id=$1", id)

	err := row.Scan(&result.Id, &result.Title, &result.Author, &result.Forum, &result.Message, &result.Votes,
		&result.Slug, &result.Created)
	if err != nil {
		return models.Thread{}, err
	}

	return result, nil
}