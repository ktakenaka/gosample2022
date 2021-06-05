package logger

import (
	"context"
)

// Logger is interface to log application info
type Logger interface {
	// Error prints error log
	Error(ctx context.Context, msg string)
	// Errorf prints error log with the specified msg format
	Errorf(ctx context.Context, format string, args ...interface{})

	// Info prints info log
	Info(ctx context.Context, msg string)
	// Infof prints info log with the specified msg format
	Infof(ctx context.Context, format string, args ...interface{})

	// Warn prints warn log
	Warn(ctx context.Context, msg string)
	// Warnf prints warn log with the specified msg format
	Warnf(ctx context.Context, format string, args ...interface{})

	// IsDebugEnabled checks debug is possible
	IsDebugEnabled(ctx context.Context) bool
	// Debug prints debug log
	Debug(ctx context.Context, msg string)
	// Debugf prints debug log with the specified msg format
	Debugf(ctx context.Context, format string, args ...interface{})
}

type Closer interface {
	Close(ctx context.Context)
}

// CloseLogger some of loggers require close, such as Rollbar having a client
func CloseLogger(ctx context.Context, log Logger) {
	if c, ok := log.(Closer); ok {
		c.Close(ctx)
	}
}
