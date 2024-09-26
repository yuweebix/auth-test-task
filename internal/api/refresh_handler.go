package api

import (
	"encoding/json"
	"net/http"

	"github.com/yuweebix/auth-test-task/internal/models"
)

type RefreshHandler struct {
	domain domain
}

func NewRefreshHandler(domain domain) *RefreshHandler {
	return &RefreshHandler{
		domain: domain,
	}
}

func (handler *RefreshHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var req models.TokenRefreshRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := handler.domain.RefreshToken(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
