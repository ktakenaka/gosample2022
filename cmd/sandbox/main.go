package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/ktakenaka/gosample2022/app/domain/models"
	"github.com/ktakenaka/gosample2022/cmd/internal/config"
	"github.com/ktakenaka/gosample2022/infra/database"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func init() {
	boil.DebugMode = true
	boil.SetLocation(time.FixedZone("Asia/Tokyo", 9*60*60))
}

func dummyReq(i int) result {
	fmt.Println("request", i)
	time.Sleep(1 * time.Second)
	return result{res: fmt.Sprintf("ok %d", i)}
}

type result struct {
	err error
	res string
}

func main() {
	results, done := make(chan result), make(chan struct{})
	defer close(done)

	limiter := make(chan struct{}, 5)
	defer close(limiter)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(j int) {
			limiter <- struct{}{}

			defer wg.Done()
			defer func() { <-limiter }()
			for {
				select {
				case <-done:
					fmt.Printf("done %d\n", j)
					return
				case results <- dummyReq(j):
					return
				}
			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for res := range results {
		if res.err != nil {
			fmt.Println("see you")
			return
		}
		fmt.Println(res.res)
	}
}

func sampleDB() {
	cfg, err := config.Initialize()
	if err != nil {
		panic(err)
	}

	db, err := database.New(cfg.DB.Write)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	sample, _ := models.Samples().One(ctx, db)
	fmt.Println(sample.CreatedAt)
}
