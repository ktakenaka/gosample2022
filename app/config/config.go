package config

import (
	"fmt"

	"github.com/jinzhu/configor"
	"github.com/ktakenaka/gomsx/app/pkg/sqls"
)

type Config struct {
	Env string  `required:"true" env:"ENV"`
	App *AppCnf `required:"true" yaml:"app"`
	DB  *DB     `required:"true" yaml:"db"`
}

type AppCnf struct {
	ServiceName string `required:"true" yaml:"service_name" default:"gomsx"`
	Port        uint   `required:"true" yaml:"port" default:"8080"`
}

type DB struct {
	Name     string `env:"DB_NAME"`
	User     string `env:"DB_USER"`
	Host     string `env:"DB_HOST"`
	Password string `env:"DB_PASSWORD"`
	Port     uint   `env:"DB_PORT"`
}

func LoadConfig(path string) (*Config, error) {
	appConfig := Config{}
	err := configor.Load(&appConfig, path)
	if err != nil {
		return nil, fmt.Errorf("failed to load config %s: %w", path, err)
	}
	return &appConfig, nil
}

func (d *DB) SqlsConf() *sqls.Config {
	return &sqls.Config{
		Driver:       "mysql",
		DBName:       d.Name,
		Username:     d.User,
		Password:     d.Password,
		Host:         d.Host,
		Port:         d.Port,
		QueryOptions: map[string]string{"parseTime": "true"},
	}
}
