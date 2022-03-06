package config

import (
	"flag"

	"github.com/ktakenaka/gosample2022/app/config"
)

func Initialize() (*config.Config, error) {
	configFilePath := flag.String("c", "", "config file path for app")
	flag.Parse()
	if *configFilePath == "" {
		*configFilePath = "environment/local.yml"
	}

	return config.New(*configFilePath)
}
