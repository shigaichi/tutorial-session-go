package logger

import (
	"go.uber.org/zap"
)

var Logger *zap.SugaredLogger

func init() {
	conf := zap.AddCallerSkip(1)
	logger, _ := zap.NewDevelopment(conf)
	sugar := logger.Sugar()
	Logger = sugar
}

func Debug(args ...interface{}) {
	Logger.Debug(args...)
}

func Debugw(format string, args ...interface{}) {
	Logger.Debugw(format, args...)
}

func Info(args ...interface{}) {
	Logger.Info(args...)
}

func Infow(format string, args ...interface{}) {
	Logger.Infow(format, args...)
}

func Error(args ...interface{}) {
	Logger.Error(args...)
}

func Errorw(format string, args ...interface{}) {
	Logger.Errorw(format, args...)
}

func Warn(args ...interface{}) {
	Logger.Warn(args...)
}

func Warnw(format string, args ...interface{}) {
	Logger.Warnw(format, args...)
}

func Fatal(args ...interface{}) {
	Logger.Fatal(args...)
}

func Fatalw(format string, args ...interface{}) {
	Logger.Fatalw(format, args...)
}
