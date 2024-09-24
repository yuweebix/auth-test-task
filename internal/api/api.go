package api

import (
	"context"

	"github.com/yuweebix/auth-test-task/internal/models"
)

type domain interface {
	AccessToken(context.Context, *models.TokenAccessRequest) (*models.TokenAccessResponse, error)
	RefreshToken(context.Context, *models.TokenRefreshRequest) (*models.TokenRefreshResponse, error)
}

type API struct {
	domain
}

func NewAPI(domain domain) (api *API) {
	api = &API{
		domain: domain,
	}
	return
}
