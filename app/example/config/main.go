package main

import (
	"fmt"
	"log"

	"github.com/ktakenaka/gomsx/app/config"
)

func main() {
	cnf, err := config.LoadConfig("config/local.yml")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("config: %#v", cnf)
}
