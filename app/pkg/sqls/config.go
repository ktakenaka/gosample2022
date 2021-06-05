package sqls

import (
	"fmt"
	"strings"
	"time"
)

type Config struct {
	Driver string

	Username string
	Password string
	Host     string
	Port     uint
	DBName   string

	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration

	QueryOptions map[string]string
}

func (c *Config) ToConnString() string {
	dst := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		c.Username, c.Password, c.Host, c.Port, c.DBName,
	)

	if c.QueryOptions != nil {
		var options []string
		for k, v := range c.QueryOptions {
			options = append(options, k+"="+v)
		}
		dst = dst + "?" + strings.Join(options, "&")
	}
	return dst
}
