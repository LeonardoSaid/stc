package commands

import (
	"context"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/leonardosaid/stc/accounts/pkg/db"

	v1 "github.com/leonardosaid/stc/accounts/internal/api/v1"
	"github.com/leonardosaid/stc/accounts/pkg/logger"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewServerCommand(ctx context.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "server",
		Short: "Start server",
		Long:  `TODO`,
		Args:  cobra.MinimumNArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			fx.New(
				fx.Provide(
					logger.NewZapLogger,
					db.ConnectPostgres,
					newEcho,
				),
				v1.Module,
				fx.Invoke(serverHandler),
			).Run()
			return nil
		},
	}
}

func newEcho(log *zap.Logger) *echo.Echo {
	e := echo.New()
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	g := e.Group("/healthcheck")
	g.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		LogErrorFunc: func(c echo.Context, err error, stack []byte) error {
			log.Error("HTTP Server recovered from panic", zap.Error(err))
			return err
		},
	}))
	e.HTTPErrorHandler = v1.DefaultHTTPErrorHandler(log)
	return e
}

func serverHandler(lc fx.Lifecycle, e *echo.Echo, log *zap.Logger) *echo.Echo {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			log.Info("Starting HTTP server")
			go func() {
				address := os.Getenv("SERVER_ADDRESS")
				if err := e.Start(address); !errors.Is(err, http.ErrServerClosed) {
					log.Fatal("Error running server", zap.Error(err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info("Stopping HTTP server")
			cancelCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
			defer cancel()
			if err := e.Shutdown(cancelCtx); err != nil {
				log.Error("Error shutting down server", zap.Error(err))
			} else {
				log.Info("Server shutdown gracefully")
			}
			return log.Sync()
		},
	})
	return e
}
