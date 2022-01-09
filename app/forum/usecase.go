package forum

import (
	"github.com/Lyalyashechka/bdProject/app/models"
	"github.com/Lyalyashechka/bdProject/app/tools"
)

type UseCase interface {
	CreateForum(forum models.Forum) (models.Forum, *models.CustomError)
	GetDetailsForum(slug string) (models.Forum, *models.CustomError)
	CreateThread(thread models.Thread) (models.Thread, *models.CustomError)
	GetUsersForum(slug string, filter tools.FilterUser) ([]models.User, *models.CustomError)
	GetForumThreads(slug string, filter tools.FilterThread) ([]models.Thread, *models.CustomError)
}
