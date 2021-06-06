package httpserver

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ktakenaka/gomsx/app/cmd/db"
	"github.com/ktakenaka/gomsx/app/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/ktakenaka/gomsx/app/internal/models/v1.0/dao/mysql"
	"github.com/ktakenaka/gomsx/app/internal/rest/v1.0/handlers/sample"
)

type Task struct {
	server *echo.Echo
}

func Initialize(ctx context.Context, conf *config.AppCnf, dbTask *db.Task) (*Task, error) {
	// Server
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	txmFact := mysql.TxManagerFactory(dbTask.GetDB())
	sampleHdl := sample.NewSampleHandler(ctx, txmFact)
	e.GET("/samples", sampleHdl.List)
	e.GET("/", hello)

	go func() {
		if err := e.Start(fmt.Sprintf(":%d", conf.Port)); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal(fmt.Errorf("server failure: %w", err))
		}
	}()

	return &Task{e}, nil
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func (t *Task) Name() string {
	return "http server"
}

func (t *Task) Shutdown(ctx context.Context) error {
	return t.server.Shutdown(ctx)
}
