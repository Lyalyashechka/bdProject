package thread

import "github.com/Lyalyashechka/bdProject/app/models"

type UseCase interface {
	CreatePosts(slugOrId string, post []models.Post) ([]models.Post, *models.CustomError)
	CreateVote(slugOrId string, vote models.Vote) (models.Thread, *models.CustomError)
	GetThreadDetails(slugOrId string) (models.Thread, *models.CustomError)
}
