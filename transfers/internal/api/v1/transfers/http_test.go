package transfers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/leonardosaid/stc/accounts/internal/domain"
	"github.com/leonardosaid/stc/accounts/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestTransfersHandler(t *testing.T) {
	t.Run("list transfers with error not found", func(t *testing.T) {
		e := echo.New()
		transferUC := new(mocks.TransferUseCase)

		transferUC.On("ListByAccountID", mock.Anything, mock.Anything).Return(nil, &domain.NotFoundError{})

		req := httptest.NewRequest(http.MethodGet, "/transfer-management/v1/transfers", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.Set("user", &jwt.Token{Claims: &domain.LoginToken{ID: uuid.NewString()}})
		h := NewTransferHandler(transferUC)

		err := h.List(c)

		assert.Error(t, err)
		assert.IsType(t, err, &domain.NotFoundError{})
	})

	t.Run("create transfer successfully", func(t *testing.T) {
		e := echo.New()
		transferUC := new(mocks.TransferUseCase)
		jsonData := `{"account_destination_id":"179bf3d9-4d18-4f64-a1e3-a206f4c0aa35","amount":1000}`

		transferUC.On("Create", mock.Anything, mock.Anything, mock.Anything).Return(nil)

		req := httptest.NewRequest(http.MethodPost, "/transfer-management/v1/transfers", strings.NewReader(jsonData))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.Set("user", &jwt.Token{Claims: &domain.LoginToken{ID: uuid.NewString()}})
		h := NewTransferHandler(transferUC)

		err := h.Create(c)

		assert.NoError(t, err)
		transferUC.AssertNumberOfCalls(t, "Create", 1)
	})
}
