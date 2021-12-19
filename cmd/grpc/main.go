package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	samplePb "github.com/ktakenaka/gosample2022/app/interface/grpc/protos/sample"
	"google.golang.org/grpc"
)

const (
	portNum = 8080
)

var (
	port = flag.Int("port", portNum, "The server port")
)

type server struct {
	samplePb.UnimplementedSampleServer
}

// LisaSamples implements samplePb.LisaSamples
func (s *server) ListSamples(ctx context.Context, in *samplePb.Request) (*samplePb.Response, error) {
	return &samplePb.Response{}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		panic(err)
	}
	log.Printf("server listening at %v", lis.Addr())

	s := grpc.NewServer()
	samplePb.RegisterSampleServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
