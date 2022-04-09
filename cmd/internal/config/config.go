package config

import (
	"os"

	"github.com/ktakenaka/gosample2022/app/config"
	"github.com/ktakenaka/gosample2022/environment"
)

func Initialize() (*config.Config, error) {
	file, err := environment.GetConfig(os.Getenv("ENV"))
	if err != nil {
		panic(err)
	}

	return config.New(file)
}
