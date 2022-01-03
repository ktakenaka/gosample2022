package config

import (
	"os"

	"gopkg.in/yaml.v3"

	"github.com/ktakenaka/gosample2022/infra/database"
	"github.com/ktakenaka/gosample2022/infra/rollbar"
)

type Config struct {
	Env string
	App struct {
		ServiceName string `yaml:"service_name"`
		Port        int
		API         string `yaml:"api"`
		IsRollbar   bool
	}
	DB struct {
		Write *database.Config
		Read  *database.Config
	}
	Rollbar *rollbar.Config
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
