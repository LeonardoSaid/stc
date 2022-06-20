package transfers

import (
	"github.com/golang-jwt/jwt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/leonardosaid/stc/accounts/internal/usecases"
)

type Handler struct {
	UseCase usecases.TransferUseCase
}

func NewTransferHandler(a usecases.TransferUseCase) Handler {
	return Handler{a}
}

func (h *Handler) List(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)

	rows, err := h.UseCase.ListByAccountID(c.Request().Context(), token)
	if err != nil {
		return err
	}

	var out []ResponseDTO
	for i := range rows {
		if m, err := ToTransferDTO(&rows[i]); err == nil {
			out = append(out, m)
		} else {
			return err
		}
	}

	return c.JSON(http.StatusOK, out)
}

func (h *Handler) Create(c echo.Context) error {
	in := new(CreateTransferDTO)
	token := c.Get("user").(*jwt.Token)

	if err := c.Bind(in); err != nil {
		return err
	}

	entity, err := in.ToEntity()
	if err != nil {
		return err
	}

	err = h.UseCase.Create(c.Request().Context(), &entity, token)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
