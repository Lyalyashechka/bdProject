package forum

import (
	"github.com/Lyalyashechka/bdProject/app/models"
	"github.com/Lyalyashechka/bdProject/app/tools"
)

type Repository interface {
	AddForum(forum models.Forum) (models.Forum, error)
	GetForumBySlug(slug string) (models.Forum, error)
	GetDetailsForum(slug string) (models.Forum, error)
	AddThread(thread models.Thread) (models.Thread, error)
	GetUsersForum(slug string, filter tools.FilterUser) ([]models.User, error)
	GetForumThreads(slug string, filter tools.FilterThread) ([]models.Thread, error)
}
