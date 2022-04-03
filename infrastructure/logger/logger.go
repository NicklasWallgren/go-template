package logger

import (
	"fmt"
	"github.com/NicklasWallgren/go-template/infrastructure/env"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	gormLogger "gorm.io/gorm/logger"
	"moul.io/zapgorm2"
)

// Logger structure
type Logger struct {
	*zap.SugaredLogger
	zapLogger *zap.Logger
}

type GinLogger struct {
	*Logger
}

// Write interface implementation for gin-framework
func (l GinLogger) Write(p []byte) (n int, err error) {
	l.Info(string(p)) // TODO
	return len(p), nil
}

// NewLogger get the logger
func NewLogger(env env.Env) (Logger, error) {
	level, err := zapcore.ParseLevel(env.LogLevel)
	if err != nil {
		return Logger{}, fmt.Errorf("unable to determine log level %w", err)
	}

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.Level.SetLevel(level)

	zapLogger, err := config.Build()
	if err != nil {
		return Logger{}, fmt.Errorf("unable to create logger %w", err)
	}

	return Logger{SugaredLogger: zapLogger.Sugar(), zapLogger: zapLogger}, nil
}

// GetGinLogger get the gin logger
func (l Logger) GetGinLogger() GinLogger {
	logger := l.zapLogger.WithOptions(zap.WithCaller(false))

	return GinLogger{Logger: &Logger{SugaredLogger: logger.Sugar()}}
}

// GetFxLogger gets logger for go-fx
func (l *Logger) GetFxLogger() fxevent.Logger {
	logger := l.zapLogger.WithOptions(zap.WithCaller(false))

	return &fxevent.ZapLogger{Logger: logger}
}

// GetGormLogger gets the gorm framework logger
func (l Logger) GetGormLogger() gormLogger.Interface {
	log := zapgorm2.New(l.zapLogger)

	return log.LogMode(gormLogger.Info)
}
