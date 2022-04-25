package logger

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
	//	"go.uber.org/zap/zapcore"
)

func NewLogger() *zap.SugaredLogger {
	logger, _ := zap.NewProduction()
	return logger.Sugar()
}

// Module provided to fx
var Module = fx.Options(
	fx.Provide(NewLogger),
)

// Logger is a simplified abstraction of the zap.Logger
// type ILogger interface {
// 	Debug(msg string, fields ...zapcore.Field)
// 	Info(msg string, fields ...zapcore.Field)
// 	Warn(msg string, fields ...zapcore.Field)
// 	Error(msg string, fields ...zapcore.Field)
// 	Fatal(msg string, fields ...zapcore.Field)
// 	With(fields ...zapcore.Field) Logger
// }

// type Logger struct {
// 	logger *zap.Logger
// }

// func (l Logger) Debug(msg string, fields ...zapcore.Field) {
// 	l.logger.Debug(msg, fields...)
// }

// func (l Logger) Info(msg string, fields ...zapcore.Field) {
// 	l.logger.Info(msg, fields...)
// }

// func (l Logger) Warn(msg string, fields ...zapcore.Field) {
// 	l.logger.Warn(msg, fields...)
// }

// func (l Logger) Error(msg string, fields ...zapcore.Field) {
// 	l.logger.Error(msg, fields...)
// }

// func (l Logger) Fatal(msg string, fields ...zapcore.Field) {
// 	l.logger.Fatal(msg, fields...)
// }

// func (l Logger) With(fields ...zapcore.Field) Logger {
// 	return Logger{logger: l.logger.With(fields...)}
// }
