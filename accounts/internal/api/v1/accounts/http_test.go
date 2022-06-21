package accounts

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/leonardosaid/stc/accounts/internal/domain"
	"github.com/leonardosaid/stc/accounts/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAccountsHandler(t *testing.T) {
	t.Run("create account successfully", func(t *testing.T) {
		e := echo.New()
		accountUC := new(mocks.AccountUseCase)
		jsonData := `{"name":"Fulano","cpf":"00000000191","secret":"secret"}`

		accountUC.On("Create", mock.Anything, mock.Anything).Return(nil)

		req := httptest.NewRequest(http.MethodPost, "/account-management/v1/accounts", strings.NewReader(jsonData))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := NewAccountHandler(accountUC)

		if assert.NoError(t, h.Create(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
		}
	})

	t.Run("list accounts with error not found", func(t *testing.T) {
		e := echo.New()
		accountUC := new(mocks.AccountUseCase)

		accountUC.On("List", mock.Anything, mock.Anything).Return(nil, &domain.NotFoundError{})

		req := httptest.NewRequest(http.MethodGet, "/account-management/v1/accounts", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := NewAccountHandler(accountUC)

		err := h.List(c)

		assert.Error(t, err)
		assert.IsType(t, err, &domain.NotFoundError{})
	})
}
