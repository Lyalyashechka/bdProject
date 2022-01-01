package handler

import (
	"github.com/Lyalyashechka/bdProject/app/models"
	"github.com/Lyalyashechka/bdProject/app/user"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	UseCase user.UseCase
}

func NewHandler(useCase user.UseCase) *Handler {
	return &Handler{
		UseCase: useCase,
	}
}

func (handler *Handler) SignUpUser(ctx echo.Context) error {
	var newUser models.User

	if err := ctx.Bind(&newUser); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	if err := ctx.Validate(&newUser); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	newUser.Nickname = ctx.Param("nickname")
	users, err := handler.UseCase.CreateUser(newUser)
	if err != nil {
		if users[0].Email == "" {
			return ctx.JSON(http.StatusInternalServerError, err.Error())
		}
		return ctx.JSON(http.StatusConflict, users)
	}

	return ctx.JSON(http.StatusCreated, users[0])
}
