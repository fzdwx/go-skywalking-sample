package logx

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"go.uber.org/zap"
)

// ZapLogger is a logger impl.
type ZapLogger struct {
	log  *zap.Logger
	Sync func() error
}

type Logger interface {
	GetLog() *zap.Logger
	Log(level log.Level, keyvals ...interface{}) error
}

var Zap *zap.Logger

// NewZapLogger return a zap logger.
func NewZapLogger(opts ...zap.Option) *ZapLogger {
	l, err := zap.NewProduction(opts...)
	if err != nil {
		panic(err)
	}
	Zap = l
	return &ZapLogger{log: l, Sync: l.Sync}
}

func (l *ZapLogger) GetLog() *zap.Logger {
	return l.log
}

// Log Implementation of logger interface.
func (l *ZapLogger) Log(level log.Level, keyvals ...interface{}) error {
	if len(keyvals) == 0 || len(keyvals)%2 != 0 {
		l.log.Warn(fmt.Sprint("Keyvalues must appear in pairs: ", keyvals))
		return nil
	}
	// Zap.Field is used when keyvals pairs appear
	var data []zap.Field
	for i := 0; i < len(keyvals); i += 2 {
		data = append(data, zap.Any(fmt.Sprint(keyvals[i]), fmt.Sprint(keyvals[i+1])))
	}
	switch level {
	case log.LevelDebug:
		l.log.Debug("", data...)
	case log.LevelInfo:
		l.log.Info("", data...)
	case log.LevelWarn:
		l.log.Warn("", data...)
	case log.LevelError:
		l.log.Error("", data...)
	}
	return nil
}
