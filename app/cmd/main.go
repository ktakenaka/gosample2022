package main

import (
	"context"
	"flag"
	"log"

	"github.com/ktakenaka/gomsx/app/cmd/db"
	"github.com/ktakenaka/gomsx/app/cmd/httpserver"
	"github.com/ktakenaka/gomsx/app/cmd/logger"
	"github.com/ktakenaka/gomsx/app/cmd/shutdown"
	"github.com/ktakenaka/gomsx/app/config"
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

	serverTask, err := httpserver.Initialize(ctx, conf.App, dbTask)
	if err != nil {
		log.Fatal(err)
	}
	shutdownTaks.Add(serverTask)

	shutdownTaks.WaitForServerStop(ctx)
}
