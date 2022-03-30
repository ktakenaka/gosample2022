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

type st struct {
	F1 string
	F2 int
	F3 []byte
	F4 bool
}

var _ encoding.BinaryMarshaler = (*st)(nil)
var _ encoding.BinaryUnmarshaler = (*st)(nil)

func (s *st) MarshalBinary() (data []byte, err error) {
	return json.Marshal(s)
}

func (s *st) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, s)
}

type map4redis map[string][]int

var _ encoding.BinaryMarshaler = (*map4redis)(nil)
var _ encoding.BinaryUnmarshaler = (*map4redis)(nil)

func (m *map4redis) MarshalBinary() (data []byte, err error) {
	return json.Marshal(m)
}
func (m *map4redis) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, m)
}

func main() {
	client, task, _ := redis.Init(ctx, cfg)
	defer task.Shutdown(ctx)

	fmt.Println("struct to save and load from Redis")
	s1 := &st{F1: "hello", F2: 2, F3: []byte("see you"), F4: true}
	fmt.Println(client.Set(ctx, key, s1, 0).Result())

	s2 := &st{}
	fmt.Println(client.Get(ctx, key).Scan(s2))
	fmt.Printf("%v\n", s2)

	fmt.Println("map to save and load from Redis")
	m1 := &map4redis{"f1": []int{1}, "f2": []int{2, 3}}
	fmt.Println(client.Set(ctx, key, m1, 0).Result())

	m2 := &map4redis{}
	fmt.Println(client.Get(ctx, key).Scan(m2))
	fmt.Printf("%v\n", m2)
}
