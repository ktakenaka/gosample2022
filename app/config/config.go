package config

import (
	"github.com/friendsofgo/errors"
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
	return &appConfig, errors.Wrapf(err, "failed to load config %s", path)
}

func (d *DB) SqlsConf() *sqls.Config {
	return &sqls.Config{
		Driver:   "mysql",
		Username: d.Name,
		Password: d.Password,
		Host:     d.Host,
		Port:     d.Port,
	}
}
