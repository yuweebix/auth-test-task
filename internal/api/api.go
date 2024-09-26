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

func RegisterServeMux(domain domain) (mux *http.ServeMux) {
	mux = http.NewServeMux()

	// Tokens
	mux.Handle("POST /access", NewAccessHandler(domain))
	mux.Handle("POST /refresh", NewRefreshHandler(domain))

	return
}
