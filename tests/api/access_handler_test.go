package api_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/yuweebix/auth-test-task/internal/api"
	mocks "github.com/yuweebix/auth-test-task/mocks/api"
)

func TestAccessNoUserID(t *testing.T) {
	r := httptest.NewRequest("POST", "/access", nil)
	w := httptest.NewRecorder()

	domain := mocks.NewMockDomain(t)
	handler := api.NewAccessHandler(domain)

	handler.ServeHTTP(w, r)

	res := w.Result()
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	require.NoError(t, err)

	assert.Equal(t, res.StatusCode, http.StatusBadRequest)
	assert.Contains(t, string(body), "missing user_id")
}

func TestAccessEmptyUserID(t *testing.T) {
	r := httptest.NewRequest("POST", "/access?user_id=", nil)
	w := httptest.NewRecorder()

	domain := mocks.NewMockDomain(t)
	handler := api.NewAccessHandler(domain)

	handler.ServeHTTP(w, r)

	res := w.Result()
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	require.NoError(t, err)

	assert.Equal(t, res.StatusCode, http.StatusBadRequest)
	assert.Contains(t, string(body), "missing user_id")
}

func TestAccessInvalidUserID(t *testing.T) {
	r := httptest.NewRequest("POST", "/access?user_id=123", nil)
	w := httptest.NewRecorder()

	domain := mocks.NewMockDomain(t)
	handler := api.NewAccessHandler(domain)

	handler.ServeHTTP(w, r)

	res := w.Result()
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	require.NoError(t, err)

	assert.Equal(t, res.StatusCode, http.StatusBadRequest)
	assert.Contains(t, string(body), "invalid UUID length")
}
