package logger

import (
	"fmt"

	"github.com/NicklasWallgren/go-template/adapters/driven/env"

	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	gormLogger "gorm.io/gorm/logger"
	"moul.io/zapgorm2"
)

type Logger interface {
	Debug(args ...interface{})
	Debugf(template string, args ...interface{})
	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Warn(args ...interface{})
	Warnf(template string, args ...interface{})
	Error(args ...interface{})
	Errorf(template string, args ...interface{})
	With(args ...interface{}) Logger
	GetFxLogger() fxevent.Logger
	GetGormLogger() gormLogger.Interface
	GetZapLogger() *zap.Logger
}

// logger structure.
type logger struct {
	*zap.SugaredLogger
	zapLogger *zap.Logger
}

// NewLogger get the logger.
func NewLogger(env env.Env) (Logger, error) {
	level, err := zapcore.ParseLevel(env.LogLevel)
	if err != nil {
		return logger{}, fmt.Errorf("unable to determine log level %w", err)
	}

	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.Level.SetLevel(level)

	zapLogger, err := config.Build()
	if err != nil {
		return logger{}, fmt.Errorf("unable to create logger %w", err)
	}

	return logger{SugaredLogger: zapLogger.Sugar(), zapLogger: zapLogger}, nil
}

func (l logger) With(fields ...interface{}) Logger {
	l.SugaredLogger = l.SugaredLogger.With(fields)

	return l
}

// GetFxLogger gets logger for go-fx.
func (l logger) GetFxLogger() fxevent.Logger {
	zapLogger := l.zapLogger.WithOptions(zap.WithCaller(false))

	return NewFxZapLoggerDecorator(zapLogger)
}

// GetGormLogger gets the gorm framework logger.
func (l logger) GetGormLogger() gormLogger.Interface {
	zapGormLogger := zapgorm2.New(l.zapLogger)

	return zapGormLogger.LogMode(gormLogger.Info)
}

// GetZapLogger gets the gorm framework logger.
func (l logger) GetZapLogger() *zap.Logger {
	return l.zapLogger
}
