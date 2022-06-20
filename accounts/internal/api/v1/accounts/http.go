package accounts

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/leonardosaid/stc/accounts/internal/domain"
	"github.com/leonardosaid/stc/accounts/internal/usecases"
	"net/http"
)

type Handler struct {
	UseCase usecases.AccountUseCase
}

func NewAccountHandler(a usecases.AccountUseCase) Handler {
	return Handler{a}
}

func (h *Handler) List(c echo.Context) error {
	rows, err := h.UseCase.List(c.Request().Context())
	if err != nil {
		return err
	}

	var out []ResponseDTO
	for i := range rows {
		if m, err := ToAccountDTO(&rows[i]); err == nil {
			out = append(out, m)
		} else {
			return err
		}
	}

	return c.JSON(http.StatusOK, out)
}

func (h *Handler) Create(c echo.Context) error {
	in := new(CreateAccountDTO)

	if err := c.Bind(in); err != nil {
		return err
	}

	entity, err := in.ToEntity()
	if err != nil {
		return err
	}

	err = h.UseCase.Create(c.Request().Context(), &entity)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *Handler) FindBalanceByID(c echo.Context) error {
	var in string
	var id uuid.UUID

	if err := echo.PathParamsBinder(c).String("id", &in).BindError(); err != nil {
		return err
	}

	id, err := uuid.Parse(in)
	if err != nil {
		return err
	}

	amount, err := h.UseCase.FindBalanceByID(c.Request().Context(), id)
	if err != nil {
		return err
	}

	out := ToBalanceDTO(amount)

	return c.JSON(http.StatusOK, out)
}

func (h *Handler) FindByCPF(c echo.Context) error {
	var cpf string

	if err := echo.PathParamsBinder(c).String("id", &cpf).BindError(); err != nil {
		return err
	}

	acc, err := h.UseCase.FindByCPF(c.Request().Context(), cpf)
	if err != nil {
		return err
	}

	out, err := ToAccountDTO(acc)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, out)
}

func (h *Handler) UpdateBalanceByID(c echo.Context) error {
	in := new(UpdateBalanceDTO)

	if err := c.Bind(in); err != nil {
		return err
	}

	id, err := uuid.Parse(in.ID)
	if err != nil {
		return err
	}

	acc := &domain.Account{
		ID:      id,
		Balance: in.Balance,
	}

	err = h.UseCase.UpdateBalanceByID(c.Request().Context(), acc)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}
