package thread

import "github.com/Lyalyashechka/bdProject/app/models"

type Repository interface {
	CreatePosts(threadId int, threadForum string, post []models.Post) ([]models.Post, error)
	GetThreadBySlug(slug string) (models.Thread, error)
	GetThreadById(id int) (models.Thread, error)
	GetThreadBySlugOrId(slugOrId string) (models.Thread, error)
	CreateVoteBySlugOrId(slugOrId string, vote models.Vote) error
	UpdateVoteBySlugOrId(slugOrId string, vote models.Vote) error
}
