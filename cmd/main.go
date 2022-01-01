package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/ktakenaka/gosample2022/cmd/config"
	"github.com/ktakenaka/gosample2022/cmd/database"
	"github.com/ktakenaka/gosample2022/cmd/grpc"
	"github.com/ktakenaka/gosample2022/cmd/shutdown"
)

func main() {
	cfg, err := config.Initialize()
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tasks := shutdown.New()
	defer tasks.Shutdown(ctx)

	read, write, task, err := database.Init(cfg.DB.Write, cfg.DB.Read)
	if err != nil {
		panic(err)
	}
	tasks.Add(task)

	s, task := grpc.New(read, write)
	tasks.Add(task)

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", cfg.App.Port))
	if err != nil {
		panic(err)
	}
	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
