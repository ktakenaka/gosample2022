package config

import (
	"errors"
	"flag"

	"github.com/ktakenaka/gosample2022/app/config"
)

func Initialize() (*config.Config, error) {
	configFilePath := flag.String("c", "", "config file path for app")
	flag.Parse()
	if *configFilePath == "" {
		return nil, errors.New("Not having config file")
	}

	return config.New(*configFilePath)
}
