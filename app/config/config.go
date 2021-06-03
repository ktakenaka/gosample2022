package config

import (
	"github.com/friendsofgo/errors"
	"github.com/jinzhu/configor"
)

type Config struct {
	Env string `required:"true" env:"ENV"`

	DB struct {
		Name     string `env:"DB_NAME"`
		User     string `env:"DB_USER"`
		Password string `env:"DB_PASSWORD"`
		Port     uint   `env:"DB_PORT"`
	}
}

func LoadConfig(path string) (*Config, error) {
	appConfig := Config{}
	err := configor.Load(&appConfig, path)
	return &appConfig, errors.Wrapf(err, "failed to load config %s", path)
}
