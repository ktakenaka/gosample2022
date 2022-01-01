package config

import (
	"errors"
	"flag"

	appCfg "github.com/ktakenaka/gosample2022/app/config"
)

func Initialize() (*appCfg.Config, error) {
	configFilePath := flag.String("c", "", "config file path for app")
	flag.Parse()
	if configFilePath == nil {
		return nil, errors.New("Not having config file")
	}

	return appCfg.New(*configFilePath)
}
