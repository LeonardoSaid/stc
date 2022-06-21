package v1

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/leonardosaid/stc/accounts/internal/api/v1/transfers"
	"github.com/leonardosaid/stc/accounts/internal/clients"
	"github.com/leonardosaid/stc/accounts/internal/domain"
	"github.com/leonardosaid/stc/accounts/internal/infra/repositories"
	"github.com/leonardosaid/stc/accounts/internal/usecases"
	"go.uber.org/fx"
)

func routes(e *echo.Echo, t transfers.Handler) {
	g := e.Group("transfer-management/v1")

	config := middleware.JWTConfig{
		Claims:     &domain.LoginToken{},
		SigningKey: []byte(os.Getenv("JWT_TOKEN_SECRET")),
	}
	g.Use(middleware.JWTWithConfig(config))

	g.GET("/transfers", t.List)
	g.POST("/transfers", t.Create)
}

var Module = fx.Options(
	fx.Provide(
		clients.NewServiceClientsImpl,
		transfers.NewTransferHandler,
		usecases.NewTransferUseCaseImpl,
		repositories.NewTransferRepositoryImpl,
	),
	fx.Invoke(routes),
)
