package logger

import (
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

// FxZapLoggerDecorator is a Fx event logger decorator that logs events to Zap.
type FxZapLoggerDecorator struct {
	FxZapLogger fxevent.ZapLogger
}

var _ fxevent.Logger = (*FxZapLoggerDecorator)(nil)

func NewFxZapLoggerDecorator(zapLogger *zap.Logger) *FxZapLoggerDecorator {
	return &FxZapLoggerDecorator{FxZapLogger: fxevent.ZapLogger{
		Logger: zapLogger,
	}}
}

// LogEvent logs the given event to the provided Zap logger.
func (l *FxZapLoggerDecorator) LogEvent(event fxevent.Event) {
	switch e := event.(type) {
	case *fxevent.Provided:
		for _, rtype := range e.OutputTypeNames {
			l.FxZapLogger.Logger.Debug("provided", // Logs as INFO in the default fx zap logger
				zap.String("constructor", e.ConstructorName),
				moduleField(e.ModuleName),
				zap.String("type", rtype),
			)
		}
		if e.Err != nil { // nolint: wsl
			l.FxZapLogger.Logger.Error("error encountered while applying options",
				moduleField(e.ModuleName),
				zap.Error(e.Err))
		}
	case *fxevent.Invoking:
		// Do not log stack as it will make logs hard to read.
		l.FxZapLogger.Logger.Debug("invoking", // Logs as INFO in the default fx zap logger
			zap.String("function", e.FunctionName),
			moduleField(e.ModuleName),
		)
	case *fxevent.LoggerInitialized:
		if e.Err != nil {
			l.FxZapLogger.Logger.Error("custom logger initialization failed", zap.Error(e.Err))
		} else {
			// Logs as INFO in the default fx zap logger
			l.FxZapLogger.Logger.Debug("initialized custom fxevent.Logger", zap.String("function", e.ConstructorName))
		}

	default:
		l.FxZapLogger.LogEvent(event)
	}
}

func moduleField(name string) zap.Field {
	if len(name) == 0 {
		return zap.Skip()
	}
	return zap.String("module", name) // nolint: wsl, nlreturn
}
