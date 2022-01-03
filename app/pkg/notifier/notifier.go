package notifier

import (
	"context"
	"log"

	"github.com/ktakenaka/gosample2022/app/pkg/ulid"
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

func NewPersonContext(ctx context.Context, officeID, userID ulid.ULID) context.Context {
	// This application is multi-tenant application, and each office has many users.
	// That's why we use officeID for ID, and userID for Username
	return rollbar.NewPersonContext(ctx, &rollbar.Person{Id: officeID.String(), Username: userID.String()})
}

type ntfr struct{}

func NewStd() Notifier {
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
