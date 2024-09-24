package api

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"

	"github.com/yuweebix/auth-test-task/internal/models"
)

func (api API) newMux() (mux *http.ServeMux) {
	mux = http.NewServeMux()

	mux.HandleFunc("POST /access", api.accessHandler)
	mux.HandleFunc("POST /refresh", api.refreshHandler)

	return
}

func (api API) accessHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		http.Error(w, "no user ID", http.StatusBadRequest)
		return
	}

	id, err := uuid.Parse(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	req := &models.TokenAccessRequest{
		UserID: id,
	}

	resp, err := api.AccessToken(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (api API) refreshHandler(w http.ResponseWriter, r *http.Request) {
	var req models.TokenRefreshRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := api.RefreshToken(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
