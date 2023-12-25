package shared

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

func ConfigLogger() *zap.Logger {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.ConsoleSeparator = "  |  "
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.EncodeTime = CustomTimeEncoder

	logger, _ := config.Build()
	return logger
}

func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006/01/02 15:04:05"))
}
