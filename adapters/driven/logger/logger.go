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
	GetGinLogger() GinLogger
	GetFxLogger() fxevent.Logger
	GetGormLogger() gormLogger.Interface
}

// logger structure.
type logger struct {
	*zap.SugaredLogger
	zapLogger *zap.Logger
}

type GinLogger struct {
	*logger
}

// Write interface implementation for gin-framework.
func (l GinLogger) Write(p []byte) (n int, err error) {
	l.Info(string(p)) // TODO

	return len(p), nil
}

// NewLogger get the logger.
func NewLogger(env env.Env) (Logger, error) {
	level, err := zapcore.ParseLevel(env.LogLevel)
	if err != nil {
		return logger{}, fmt.Errorf("unable to determine log level %w", err)
	}

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.Level.SetLevel(level)

	zapLogger, err := config.Build()
	if err != nil {
		return logger{}, fmt.Errorf("unable to create logger %w", err)
	}

	return logger{SugaredLogger: zapLogger.Sugar(), zapLogger: zapLogger}, nil
}

// GetGinLogger get the gin logger.
func (l logger) GetGinLogger() GinLogger {
	zapLogger := l.zapLogger.WithOptions(zap.WithCaller(false))

	l.Error()

	return GinLogger{logger: &logger{SugaredLogger: zapLogger.Sugar()}}
}

// GetFxLogger gets logger for go-fx.
func (l logger) GetFxLogger() fxevent.Logger {
	zapLogger := l.zapLogger.WithOptions(zap.WithCaller(false))

	return &fxevent.ZapLogger{Logger: zapLogger}
}

// GetGormLogger gets the gorm framework logger.
func (l logger) GetGormLogger() gormLogger.Interface {
	zapGormLogger := zapgorm2.New(l.zapLogger)

	return zapGormLogger.LogMode(gormLogger.Info)
}
