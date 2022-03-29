package main

import (
	"context"
	"encoding"
	"encoding/json"
	"fmt"

	"github.com/ktakenaka/gosample2022/cmd/internal/config"
	"github.com/ktakenaka/gosample2022/cmd/internal/redis"
)

var (
	cfg, _ = config.Initialize()
	ctx    = context.Background()
)

const (
	key = "key"
)

var _ encoding.BinaryMarshaler = (*st)(nil)
var _ encoding.BinaryUnmarshaler = (*st)(nil)

type st struct {
	F1 string
	F2 int
	F3 []byte
	F4 bool
}

func (s *st) MarshalBinary() (data []byte, err error) {
	return json.Marshal(s)
}

func (s *st) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, s)
}

func main() {
	client, task, _ := redis.Init(ctx, cfg)
	defer task.Shutdown(ctx)

	s1 := &st{F1: "hello", F2: 2, F3: []byte("see you"), F4: true}
	fmt.Println(client.Set(ctx, key, s1, 0).Result())

	strCmd := client.Get(ctx, key)
	s2 := &st{}
	fmt.Println(strCmd.Scan(s2))
	fmt.Printf("%v\n", s2)
}
