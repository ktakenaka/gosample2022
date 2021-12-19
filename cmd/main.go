package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/ktakenaka/gosample2022/cmd/database"
	"github.com/ktakenaka/gosample2022/cmd/grpc"
	"github.com/ktakenaka/gosample2022/cmd/shutdown"
)

const (
	portNum = 8080
)

var (
	port = flag.Int("port", portNum, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	tasks := shutdown.New()
	defer tasks.Shutdown(ctx)

	read, write, task, err := database.Init()
	if err != nil {
		panic(err)
	}
	tasks.Add(task)

	s, task := grpc.New(read, write)
	tasks.Add(task)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
