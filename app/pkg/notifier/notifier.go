package notifier

import (
	"context"
	"fmt"
	"log"
	"regexp"

	"github.com/rollbar/rollbar-go"
)

const (
	CRIT  = rollbar.CRIT
	ERR   = rollbar.ERR
	WARN  = rollbar.WARN
	INFO  = rollbar.INFO
	DEBUG = rollbar.DEBUG
)

var (
	ntfr notifier = &localNotifier{}

	ErrNoToken = fmt.Errorf("no token")

	scrubHeaders = regexp.MustCompile("Authorization|Cookie|Token")
	scrubFields  = regexp.MustCompile("password|secret|token")
)

func NewPersonContext(ctx context.Context, officeID, userID string) context.Context {
	// This application is multi-tenant application, and each office has many users.
	// That's why we use officeID for Username.
	return rollbar.NewPersonContext(ctx, &rollbar.Person{Id: userID, Username: officeID})
}

func Message(level string, msg string) {
	ntfr.Message(level, msg)
}

func Error(err error) {
	ntfr.ErrorWithLevel(ERR, err)
}

func ErrorWithExtrasAndContext(ctx context.Context, level string, err error, extras map[string]interface{}) {
	ntfr.ErrorWithExtrasAndContext(ctx, level, err, extras)
}

func Close() error {
	return ntfr.Close()
}

type Config struct {
	Token       string
	Environment string
}

func InitRollbar(cfg *Config) error {
	if cfg.Token == "" {
		return ErrNoToken
	}

	client := rollbar.New(cfg.Token, cfg.Environment, "", "", "")
	client.SetScrubFields(scrubFields)
	client.SetScrubHeaders(scrubHeaders)

	ntfr = client
	return nil
}

type notifier interface {
	Message(level string, msg string)
	ErrorWithLevel(level string, err error)
	ErrorWithExtrasAndContext(ctx context.Context, level string, err error, extras map[string]interface{})

	Close() error
}

type localNotifier struct{}

func (n *localNotifier) ErrorWithExtrasAndContext(
	ctx context.Context,
	level string,
	err error,
	extras map[string]interface{},
) {
	log.Printf("%s\nctx: %v\nerr: %v\nextras: %v\n", level, ctx, err, extras)
}

func (n *localNotifier) Message(level string, msg string) {
	log.Printf("%s\nmsg: %s\n", level, msg)
}

func (n *localNotifier) ErrorWithLevel(level string, err error) {
	n.Message(level, err.Error())
}

func (n *localNotifier) Close() error {
	return nil
}
