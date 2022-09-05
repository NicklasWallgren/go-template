package logger

import (
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	gormLogger "gorm.io/gorm/logger"
)

type NullLogger struct {
}

func (n NullLogger) Debug(args ...interface{}) {

}

func (n NullLogger) Debugf(template string, args ...interface{}) {

}

func (n NullLogger) Info(args ...interface{}) {

}

func (n NullLogger) Infof(template string, args ...interface{}) {

}

func (n NullLogger) Warn(args ...interface{}) {

}

func (n NullLogger) Warnf(template string, args ...interface{}) {

}

func (n NullLogger) Error(args ...interface{}) {

}

func (n NullLogger) Errorf(template string, args ...interface{}) {

}

func (n NullLogger) GetFxLogger() fxevent.Logger {
	return nil
}

func (n NullLogger) GetGormLogger() gormLogger.Interface {
	return nil
}

func (n NullLogger) GetZapLogger() *zap.Logger {
	return nil
}

