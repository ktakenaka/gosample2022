package config

import (
	"os"
	"time"

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
	Redis struct {
		URL          string        `yaml:"url"`
		PoolSize     int           `yaml:"pool_size"`
		MinIdleConns int           `yaml:"min_idle_conns"`
		ReadTimeout  time.Duration `yaml:"read_timeout"`
		WriteTimeout time.Duration `yaml:"write_timeout"`
		UseTLS       bool          `yaml:"use_tls"`
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
