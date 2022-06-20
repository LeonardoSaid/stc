package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/leonardosaid/stc/accounts/internal/api/v1/login"
	"github.com/leonardosaid/stc/accounts/internal/clients"
	"github.com/leonardosaid/stc/accounts/internal/usecases"
	"go.uber.org/fx"
)

func routes(e *echo.Echo, l login.Handler) {
	g := e.Group("login-management/v1")

	g.POST("/login", l.Login)
}

var Module = fx.Options(
	fx.Provide(
		clients.NewServiceClientsImpl,
		login.NewLoginHandler,
		usecases.NewLoginUseCaseImpl,
	),
	fx.Invoke(routes),
)
