package v1

import (
	"github.com/leonardosaid/stc/accounts/internal/domain"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

const (
	InternalErrorMessage = "internal service error"
)

type ErrorMessage struct {
	Value string `json:"error"`
}

func DefaultHTTPErrorHandler(logger *zap.Logger) echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		logger.Error("Error", zap.Error(err))
		var code int
		message := &ErrorMessage{}

		switch err.(type) {
		case *domain.ValidationError:
			message.Value = err.Error()
			code = http.StatusBadRequest
		case *domain.InvalidCredentialsError:
			message.Value = err.Error()
			code = http.StatusForbidden
		case *domain.NotFoundError:
			code = http.StatusNotFound
		case domain.UnprocessableError:
			message.Value = err.Error()
			code = http.StatusUnprocessableEntity
		default:
			message.Value = InternalErrorMessage
			code = http.StatusInternalServerError
		}

		if err := c.JSON(code, message); err != nil {
			logger.Error("Error to return", zap.Error(err))
		}
	}
}
