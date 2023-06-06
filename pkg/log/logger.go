package log

import (
	"context"
	"go.uber.org/zap"
)

type Logger interface {
	With(ctx context.Context, args ...interface{}) Logger
	Info(args ...interface{})
	Debug(args ...interface{})
	Error(args ...interface{})
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}

type logger struct{ *zap.SugaredLogger }
type contextKey int

const (
	requestIDKey contextKey = iota
	correlationIDKey
)

func New() Logger {
	l, _ := zap.NewProduction()
	return NewWithZap(l)
}

func NewWithZap(l *zap.Logger) Logger {
	return &logger{l.Sugar()}
}

func (l *logger) With(ctx context.Context, args ...interface{}) Logger {
	if ctx != nil {
		if id, ok := ctx.Value(requestIDKey).(string); ok {
			args = append(args, zap.String("request_id", id))
		}

		if id, ok := ctx.Value(correlationIDKey).(string); ok {
			args = append(args, zap.String("correlation_id", id))
		}
	}

	if len(args) > 0 {
		return &logger{l.SugaredLogger.With(args...)}
	}

	return l
}
