package config

import (
	"os"

	"gopkg.in/yaml.v3"

	"github.com/ktakenaka/gosample2022/infra/database"
)

type Config struct {
	Env string
	App struct {
		ServiceName string
		Port        int
	}
	DB struct {
		Write *database.Config
		Read  *database.Config
	}
}

func New(configFilePath string) (*Config, error) {
	cfgByte, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}

	cfgByte = []byte(os.ExpandEnv(string(cfgByte)))
	cfg := &Config{}
	err = yaml.Unmarshal(cfgByte, &cfg)
	return cfg, err
}
