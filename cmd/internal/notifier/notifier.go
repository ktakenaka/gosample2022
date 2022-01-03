package notifier

import (
	"context"

	"github.com/ktakenaka/gosample2022/app/config"
	"github.com/ktakenaka/gosample2022/app/pkg/notifier"
	"github.com/ktakenaka/gosample2022/cmd/internal/shutdown"
	"github.com/ktakenaka/gosample2022/infra/rollbar"
)

type task struct {
	client interface{ Close() error }
}

func (t *task) Name() string {
	return "notifier"
}

func (t *task) Shutdown(ctx context.Context) error {
	if t.client != nil {
		return t.client.Close()
	}
	return nil
}

func Init(cfg *config.Config) (notifier.Notifier, shutdown.Task) {
	rollbarCfg := cfg.Rollbar

	if rollbarCfg.Token == "" || !cfg.App.IsRollbar {
		return notifier.NewStd(), &task{}
	}

	rollbarCfg.Environment = cfg.Env
	rollbarCfg.CodeVersion = cfg.App.API
	rollbarCfg.ServerRoot = cfg.App.ServiceName

	client := rollbar.New(rollbarCfg)
	return client, &task{client: client}
}
