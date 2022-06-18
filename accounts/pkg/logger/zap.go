package logger

import (
	"os"

	"go.uber.org/zap"
)

func NewZapLogger() (logger *zap.Logger, err error) {
	zapPreset := os.Getenv("ZAP_PRESET")
	switch {
	case zapPreset == "development":
		logger, err = zap.NewDevelopment()
	case zapPreset == "example":
		logger = zap.NewExample()
	default:
		logger, err = zap.NewProduction()
	}
	return
}
