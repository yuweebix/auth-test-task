package domain

import (
	"context"

	"github.com/yuweebix/auth-test-task/internal/schemas"
)

type repository interface {
	GetUser(context.Context, ...schemas.Filter) (*schemas.User, error)

	GetRefreshToken(context.Context, ...schemas.Filter) (*schemas.RefreshToken, error)
	DeleteRefreshToken(context.Context, ...schemas.Filter) error
}

type Domain struct {
	repository
}

func NewDomain(repository repository) (domain *Domain) {
	domain = &Domain{
		repository: repository,
	}
	return
}
