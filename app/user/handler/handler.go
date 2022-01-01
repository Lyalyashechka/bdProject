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

func (handler *Handler) GetUser(ctx echo.Context) error {
	nickname := ctx.Param("nickname")
	user, err := handler.UseCase.GetUserProfile(nickname)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, err)
	}
	return ctx.JSON(http.StatusOK, user)
}

func (handler *Handler) UpdateUser(ctx echo.Context) error {
	var UserUpdate models.User

	if err := ctx.Bind(&UserUpdate); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	if err := ctx.Validate(&UserUpdate); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	UserUpdate.Nickname = ctx.Param("nickname")

	user, err := handler.UseCase.UpdateUserProfile(UserUpdate)
	if err != nil {
		if err.Message == models.NoUser {
			return ctx.JSON(http.StatusNotFound, err)
		}
		if err.Message == models.ConflictData {
			return ctx.JSON(http.StatusConflict, err)
		}
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, user)
}
