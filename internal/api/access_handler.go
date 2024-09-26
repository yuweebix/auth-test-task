package api

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"

	"github.com/yuweebix/auth-test-task/internal/models"
)

type AccessHandler struct {
	domain domain
}

func NewAccessHandler(domain domain) *AccessHandler {
	return &AccessHandler{
		domain: domain,
	}
}

func (handler *AccessHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		http.Error(w, "missing user_id", http.StatusBadRequest)
		return
	}

	id, err := uuid.Parse(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	req := &models.TokenAccessRequest{
		UserID: id,
	}

	resp, err := handler.domain.AccessToken(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
