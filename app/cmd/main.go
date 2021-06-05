package main

import (
	"context"
	"net/http"

	"github.com/friendsofgo/errors"
	"github.com/ktakenaka/gomsx/app/cmd/shutdown"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	ctx, ctxCancel := context.WithCancel(context.Background())
	defer ctxCancel()

	shutdownTaks := shutdown.NewShutdownTasks()

	// TODO: Initialize Config
	// TODO: Connect to DB
	// TODO: Shut down task
	// TODO (Optional): logger & monitor

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
	
	shutdown.WaitForServerStop(ctx)

	shutdownTaks.ExecuteAll(ctx)
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
