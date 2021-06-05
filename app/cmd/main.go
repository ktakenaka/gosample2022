package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/friendsofgo/errors"
	"github.com/ktakenaka/gomsx/app/cmd/db"
	"github.com/ktakenaka/gomsx/app/cmd/logger"
	"github.com/ktakenaka/gomsx/app/cmd/shutdown"
	"github.com/ktakenaka/gomsx/app/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	ctx, ctxCancel := context.WithCancel(context.Background())
	defer ctxCancel()

	loggerTask := logger.Initialize(ctx)
	shutdownTaks := shutdown.NewShutdownTasks(loggerTask.AppLogger)
	defer shutdownTaks.ExecuteAll(ctx)

	shutdownTaks.Add(loggerTask)

	configFilePath := flag.String("c", "", "config file path for app")
	flag.Parse()
	if configFilePath == nil {
		log.Fatal("Please specify file path flag!")
	}
	conf, err := config.LoadConfig(*configFilePath)
	if err != nil {
		log.Fatal(err)
	}

	dbTask, err := db.Initialize(ctx, conf.DB)
	if err != nil {
		log.Fatal(err)
	}
	shutdownTaks.Add(dbTask)

	// Server
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", hello)
	go func() {
		if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal(errors.Wrap(err, "Server failure"))
		}
	}()

	shutdownTaks.WaitForServerStop(ctx)
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
