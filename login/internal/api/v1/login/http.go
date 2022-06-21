package login

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/leonardosaid/stc/accounts/internal/usecases"
)

type Handler struct {
	UseCase usecases.LoginUseCase
}

func NewLoginHandler(a usecases.LoginUseCase) Handler {
	return Handler{a}
}

func (h *Handler) Login(c echo.Context) error {
	in := new(LoginDTO)

	if err := c.Bind(in); err != nil {
		return err
	}

	entity, err := in.ToEntity()
	if err != nil {
		return err
	}

	resp, err := h.UseCase.Login(c.Request().Context(), &entity)
	if err != nil {
		return err
	}

	out := ResponseDTO{Token: resp}

	return c.JSON(http.StatusOK, out)
}
