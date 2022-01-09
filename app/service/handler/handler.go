package handler

import (
	"github.com/Lyalyashechka/bdProject/app/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	UseCase service.UseCase
}

func NewHandler (useCase service.UseCase) *Handler {
	return &Handler{UseCase: useCase}
}

func (handler *Handler) Status(ctx echo.Context) error  {
	status, err := handler.UseCase.GetStatus()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, status)
}

func (handler *Handler) Clear(ctx echo.Context) error {
	err := handler.UseCase.Clear()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.NoContent(http.StatusOK)
}
