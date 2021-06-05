package logger

import (
	"context"

	"github.com/ktakenaka/gomsx/app/pkg/logger"
	"github.com/ktakenaka/gomsx/app/pkg/logger/applogger"
)

type Task struct {
	AppLogger logger.Logger
}

func Initialize(ctx context.Context) *Task {
	// TODO: Change logger based on the environment
	// 	Like rollbar for prod
	return &Task{applogger.NewLogger()}
}

func (t *Task) Name() string {
	return "logger"
}

func (t *Task) Shutdown(ctx context.Context) error {
	logger.CloseLogger(ctx, t.AppLogger)
	return nil
}
