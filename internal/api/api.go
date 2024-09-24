package api

import (
	"context"
	"net/http"

	"github.com/yuweebix/auth-test-task/internal/models"
)

type domain interface {
	AccessToken(context.Context, *models.TokenAccessRequest) (*models.TokenAccessResponse, error)
	RefreshToken(context.Context, *models.TokenRefreshRequest) (*models.TokenRefreshResponse, error)
}

type API struct {
	domain
	server *http.Server
}

func NewAPI(domain domain) (api *API) {
	api = &API{
		domain: domain,
		server: &http.Server{
			Addr:    ":42069",
			Handler: api.newMux(),
		},
	}
	return
}

func (api API) ListenAndServe() {
	api.server.ListenAndServe()
}
