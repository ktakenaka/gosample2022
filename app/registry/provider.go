package registry

import "github.com/ktakenaka/gosample2022/app/domain/repository"

type Provider struct {
	DB    repository.DB
	Redis repository.Redis
}
