package redis

import (
	"crypto/tls"
	"time"

	"github.com/go-redis/redis/v8"
)

type Config struct {
	URL          string
	PoolSize     int
	MinIdleConns int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	UseTLS       bool
}

func New(cfg *Config) (*redis.Client, error) {
	options, err := redis.ParseURL(cfg.URL)
	if err != nil {
		return nil, err
	}

	if cfg.PoolSize > 0 {
		options.PoolSize = cfg.PoolSize
	}
	if cfg.MinIdleConns > 0 {
		options.MinIdleConns = cfg.MinIdleConns
	}
	if cfg.ReadTimeout > 0 {
		options.ReadTimeout = cfg.ReadTimeout
	}
	if cfg.WriteTimeout > 0 {
		options.WriteTimeout = cfg.WriteTimeout
	}
	if cfg.UseTLS {
		options.TLSConfig = &tls.Config{
			MinVersion: tls.VersionTLS12,
		}
	}

	client := redis.NewClient(options)
	return client, nil
}
