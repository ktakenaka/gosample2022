package repository

import (
	"github.com/go-redis/redis/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Executor interface {
	boil.ContextExecutor
}
type DB interface {
	boil.ContextExecutor
	boil.ContextBeginner
}

type Redis interface {
	redis.Cmdable
}
