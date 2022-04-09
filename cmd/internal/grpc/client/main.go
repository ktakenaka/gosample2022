package main

import (
	"context"
	"log"

	pb "github.com/ktakenaka/gosample2022/app/interface/grpc/protos/sample"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewSampleClient(conn)

	resp, err := client.SampleList(
		context.Background(),
		&pb.ListRequest{},
	)
	if err != nil {
		panic(err)
	}

	log.Printf("%v\n", resp)
}
