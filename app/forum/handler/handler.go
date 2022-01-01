package handler

import (
	"github.com/Lyalyashechka/bdProject/app/forum"
	"github.com/Lyalyashechka/bdProject/app/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	useCase forum.UseCase
}

func NewHandler(useCase forum.UseCase) *Handler {
	return &Handler{useCase: useCase}
}

func (handler *Handler) CreateForum(ctx echo.Context) error {
	var newForum models.Forum

	if err := ctx.Bind(&newForum); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	if err := ctx.Validate(&newForum); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	forum, err := handler.useCase.CreateForum(newForum)
	if err != nil {
		if err.Message == models.NoUser {
			return ctx.JSON(http.StatusNotFound, err)
		}
		if err.Message == models.ConflictData {
			return ctx.JSON(http.StatusConflict, err)
		}
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusCreated, forum)
}

func (handler *Handler) GerForumDetails(ctx echo.Context) error {
	slug := ctx.Param("slug")
	forum, err := handler.useCase.GetDetailsForum(slug)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, err)
	}
	return ctx.JSON(http.StatusOK, forum)
}
