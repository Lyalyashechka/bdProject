package forum

import "github.com/Lyalyashechka/bdProject/app/models"

type Repository interface {
	AddForum(forum models.Forum) (models.Forum, error)
	GetDetailsForum(slug string) (models.Forum, error)
}
