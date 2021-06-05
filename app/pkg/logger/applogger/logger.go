package applogger

import (
	"context"
	"log"
)

// Use just log before deciding logging strategry
type Logger struct{}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) Error(ctx context.Context, msg string) {
	log.Println(msg)
}

func (l *Logger) Errorf(ctx context.Context, format string, args ...interface{}) {
	log.Printf(format, args...)
}

func (l *Logger) Info(ctx context.Context, msg string) {
	log.Println(msg)
}

func (l *Logger) Infof(ctx context.Context, format string, args ...interface{}) {
	log.Printf(format, args...)
}

func (l *Logger) Warn(ctx context.Context, msg string) {
	log.Println(msg)
}

func (l *Logger) Warnf(ctx context.Context, format string, args ...interface{}) {
	log.Printf(format, args...)
}

func (l *Logger) IsDebugEnabled(ctx context.Context) bool {
	return true
}

func (l *Logger) Debug(ctx context.Context, msg string) {
	log.Println(msg)
}

func (l *Logger) Debugf(ctx context.Context, format string, args ...interface{}) {
	log.Printf(format, args...)
}
