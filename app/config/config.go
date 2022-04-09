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
	Mongo struct {
		URL    string `yaml:"url"`
		DBName string `yaml:"db_name"`
	}
	Kafka struct {
		Address           string `yaml:"address"`
		TopicSamples      string `yaml:"topic_samples"`
		TopicTransactions string `yaml:"topic_transactions"`
	}
	Rollbar *rollbar.Config
}

func New(cfgByte []byte) (*Config, error) {
	cfgByte = []byte(os.ExpandEnv(string(cfgByte)))
	cfg := &Config{}
	err := yaml.Unmarshal(cfgByte, &cfg)
	return cfg, err
}
