package config

import (
	"os"

	"github.com/ktakenaka/gosample2022/app/config"
	"github.com/ktakenaka/gosample2022/environment"
)

func Initialize() (*config.Config, error) {
	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}
	file, err := environment.GetConfig(env)
	if err != nil {
		return nil, err
	}

	return config.New(file)
}
