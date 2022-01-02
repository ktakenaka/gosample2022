package notifier

import (
	"context"
	"log"

	"github.com/rollbar/rollbar-go"
)

const (
	CRIT  = rollbar.CRIT
	ERR   = rollbar.ERR
	WARN  = rollbar.WARN
	INFO  = rollbar.INFO
	DEBUG = rollbar.DEBUG
)

type Notifier interface {
	Message(level string, msg string)
	ErrorWithExtrasAndContext(ctx context.Context, level string, err error, extras map[string]interface{})
}

func NewPersonContext(ctx context.Context, id string) context.Context {
	// In this application, we don't know users' email and name. That's why we just use id.
	return rollbar.NewPersonContext(ctx, &rollbar.Person{Id: id})
}

type ntfr struct{}

func New() Notifier {
	return &ntfr{}
}

func (n *ntfr) ErrorWithExtrasAndContext(
	ctx context.Context,
	level string,
	err error,
	extras map[string]interface{},
) {
	log.Printf("%s\nctx: %v\nerr: %v\nextras: %v\n", level, ctx, err, extras)
}

func (n *ntfr) Message(level string, msg string) {
	log.Printf("%s\nmsg: %s\n", level, msg)
}
