package logger

import "go.uber.org/zap"

type ZapRecoveryLogger struct {
	Logger *zap.SugaredLogger
}

func (l *ZapRecoveryLogger) Println(v ...interface{}) {
	l.Logger.Error("Recovered from panic", zap.Any("panic", v))
}
