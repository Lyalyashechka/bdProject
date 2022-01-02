package forum

import "github.com/Lyalyashechka/bdProject/app/models"

type Repository interface {
	AddForum(forum models.Forum) (models.Forum, error)
	GetDetailsForum(slug string) (models.Forum, error)
	AddThread(thread models.Thread) (models.Thread, error)
	GetUsersForum(slug string)([]models.User, error)
	GetForumThreads (slug string) ([]models.Thread, error)
}
