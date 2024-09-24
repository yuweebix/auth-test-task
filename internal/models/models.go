package models

import "github.com/google/uuid"

// Tokens

// Access

type TokenAccessRequest struct {
	UserID uuid.UUID
}

type TokenAccessResponse struct{}

// Refresh

type TokenRefreshRequest struct{}

type TokenRefreshResponse struct{}
