package rollbar

import (
	"regexp"

	"github.com/rollbar/rollbar-go"
)

type Config struct {
	Token       string
	Environment string
	Host        string
	CodeVersion string
	ServerRoot  string

	ScrubFields  *regexp.Regexp
	ScrubHeaders *regexp.Regexp
}

func New(cfg *Config) *rollbar.Client {
	client := rollbar.New(cfg.Token, cfg.Environment, cfg.CodeVersion, cfg.Host, cfg.ServerRoot)
	if cfg.ScrubFields != nil {
		client.SetScrubFields(cfg.ScrubFields)
	}
	if cfg.ScrubHeaders != nil {
		client.SetScrubHeaders(cfg.ScrubHeaders)
	}
	return client
}
