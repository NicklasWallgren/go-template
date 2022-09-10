package logger

import (
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	gormLogger "gorm.io/gorm/logger"
)

type NopLogger struct{}

func (n NopLogger) Debug(args ...interface{}) {
}

func (n NopLogger) Debugf(template string, args ...interface{}) {
}

func (n NopLogger) Info(args ...interface{}) {
}

func (n NopLogger) Infof(template string, args ...interface{}) {
}

func (n NopLogger) Warn(args ...interface{}) {
}

func (n NopLogger) Warnf(template string, args ...interface{}) {
}

func (n NopLogger) Error(args ...interface{}) {
}

func (n NopLogger) Errorf(template string, args ...interface{}) {
}

func (n NopLogger) GetFxLogger() fxevent.Logger {
	return nil
}

func (n NopLogger) GetGormLogger() gormLogger.Interface {
	return nil
}

func (n NopLogger) GetZapLogger() *zap.Logger {
	return nil
}
