package notifier

import (
	"context"

	"github.com/ktakenaka/gosample2022/app/config"
	ntfr "github.com/ktakenaka/gosample2022/app/pkg/notifier"
	"github.com/ktakenaka/gosample2022/cmd/internal/shutdown"
)

type task struct{}

func (t *task) Name() string {
	return "notifier"
}

func (t *task) Shutdown(ctx context.Context) error {
	return ntfr.Close()
}

func Init(ctx context.Context, cfg *config.Config) (shutdown.Task, error) {
	if cfg.App.IsRollbar {
		err := ntfr.InitRollbar(
			&ntfr.Config{
				Token:       cfg.Rollbar.Token,
				Environment: cfg.Env,
			},
		)
		if err != nil {
			return nil, err
		}
	}
	return &task{}, nil
}
