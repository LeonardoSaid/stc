package login

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

func TestLoginHandler(t *testing.T) {
	t.Run("login successfully", func(t *testing.T) {
		e := echo.New()
		loginUC := new(mocks.LoginUseCase)
		jsonData := `{"cpf":"00000000191","secret":"secret"}`

		loginUC.On("Login", mock.Anything, mock.Anything).Return("token", nil)

		req := httptest.NewRequest(http.MethodPost, "/login-management/v1/login", strings.NewReader(jsonData))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := NewLoginHandler(loginUC)

		if assert.NoError(t, h.Login(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	t.Run("account not found", func(t *testing.T) {
		e := echo.New()
		loginUC := new(mocks.LoginUseCase)
		jsonData := `{"cpf":"00000000191","secret":"secret"}`

		loginUC.On("Login", mock.Anything, mock.Anything).Return("", &domain.NotFoundError{})

		req := httptest.NewRequest(http.MethodPost, "/login-management/v1/login", strings.NewReader(jsonData))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		h := NewLoginHandler(loginUC)
		err := h.Login(c)

		assert.Error(t, err)
		assert.IsType(t, err, &domain.NotFoundError{})
	})
}
