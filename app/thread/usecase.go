package thread

import (
	"github.com/Lyalyashechka/bdProject/app/models"
	"github.com/Lyalyashechka/bdProject/app/tools"
)

type UseCase interface {
	CreatePosts(slugOrId string, post []models.Post) ([]models.Post, *models.CustomError)
	CreateVote(slugOrId string, vote models.Vote) (models.Thread, *models.CustomError)
	GetThreadDetails(slugOrId string) (models.Thread, *models.CustomError)
	GetPosts(slugOrId string, filter tools.FilterPosts)([]*models.Post, *models.CustomError)
	GetPost(id string, filter tools.FilterOnePost)(models.PostInfo, *models.CustomError)
	UpdateThread(slugOrId string, thread models.Thread) (models.Thread, *models.CustomError)
	UpdatePost(id string, post models.Post) (models.Post, *models.CustomError)
}
