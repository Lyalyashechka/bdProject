package forum

import "github.com/Lyalyashechka/bdProject/app/models"

type UseCase interface {
	CreateForum(forum models.Forum) (models.Forum, *models.CustomError)
	GetDetailsForum(slug string) (models.Forum, *models.CustomError)
}
