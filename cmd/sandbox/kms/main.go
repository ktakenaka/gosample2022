package main

import (
	"context"
	"fmt"

	"github.com/ktakenaka/gosample2022/cmd/internal/config"
	"github.com/ktakenaka/gosample2022/cmd/internal/kms"
)

var (
	cfg, _ = config.Initialize()
	ctx    = context.Background()
)

func main() {
	kmsConn, _ := kms.Init(ctx, cfg)
	fmt.Println(kmsConn.GenerateDataKeyPairWithoutPlaintext())
}
