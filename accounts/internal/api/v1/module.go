package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/leonardosaid/stc/accounts/internal/api/v1/accounts"
	"github.com/leonardosaid/stc/accounts/internal/infra/repositories"
	"github.com/leonardosaid/stc/accounts/internal/usecases"
	"go.uber.org/fx"
)

func routes(e *echo.Echo, a accounts.Handler) {
	g := e.Group("account-management/v1")

	g.GET("/accounts", a.List)
	g.GET("/accounts/:id", a.FindByCPF)
	g.GET("/accounts/:id/balance", a.FindBalanceByID)
	g.PATCH("/accounts/:id", a.UpdateBalanceByID)
	g.POST("/accounts", a.Create)
}

var Module = fx.Options(
	fx.Provide(
		accounts.NewAccountHandler,
		usecases.NewAccountUseCaseImpl,
		repositories.NewAccountRepositoryImpl,
	),
	fx.Invoke(routes),
)
