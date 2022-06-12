package plexapi

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewWithContextNoToken(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	client, err := NewWithContext(ctx)

	if assert.NoError(t, err, "NewWithContext should not return an error") {
		assert.Equal(t, ctx, client.ctx, "Client context should match the supplied context")
		assert.Equal(t, http.DefaultClient, client.client, "Client http.Client should be the http.DefaultClient")
	}
}

func TestNewNoToken(t *testing.T) {
	client, err := New()

	if assert.NoError(t, err, "New should not return an error") {
		assert.Equal(t, context.Background(), client.ctx, "Client context should be a background context")
		assert.Equal(t, http.DefaultClient, client.client, "Client http.Client should be the http.DefaultClient")
	}
}

func TestNewToken(t *testing.T) {
	token := "test"
	client, err := New(token)

	if assert.NoError(t, err, "New should not return an error") {
		assert.Equal(t, context.Background(), client.ctx, "Client context should be a background context")
		assert.NotEqual(t, http.DefaultClient, client.client, "Client http.Client should not be the http.DefaultClient")
		assert.IsType(t, &TokenAuthorizer{}, client.client.Transport, "Client http.Client.Transport should be a TokenAuthorizer")
		assert.Equal(t, token, client.client.Transport.(*TokenAuthorizer).token, "Client TokenAuthorizer token should match the supplied token")
	}
}

func TestClientSetTransport(t *testing.T) {
	token := "test"
	client, _ := New()
	tokenAuthorizer := NewTokenAuthorizer(token)

	client.SetTransport(tokenAuthorizer)

	assert.IsType(t, &TokenAuthorizer{}, client.client.Transport, "Client http.Client.Transport should be a TokenAuthorizer")
	assert.Equal(t, token, client.client.Transport.(*TokenAuthorizer).token, "Client TokenAuthorizer token should match the supplied token")
}

func TestNewMediaServerRequestWithContext(t *testing.T) {
	client, _ := New("test")
	client.MediaServerURL = "http://127.0.0.1"

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req, err := client.NewMediaServerRequestWithContext(ctx, http.MethodGet, "", nil)

	if assert.NoError(t, err, "NewMediaServerRequestWithContext should not return an error") {
		assert.Equal(t, ctx, req.Context(), "Context in http.Request should match the supplied context")
	}
}

func TestNewMediaServerRequestInvalidMediaServerURL(t *testing.T) {
	client, _ := New("test")
	client.MediaServerURL = string(rune(0x7f))

	_, err := client.NewMediaServerRequest(http.MethodGet, "", nil)

	assert.ErrorContains(t, err, "net/url: invalid control character in URL", "Control characters in Client.MediaServerURL should error")
}

func TestNewMediaServerRequestInvalidPath(t *testing.T) {
	client, _ := New("test")

	_, err := client.NewMediaServerRequest(http.MethodGet, string(rune(0x7f)), nil)

	assert.ErrorContains(t, err, "net/url: invalid control character in URL", "Control characters in path should error")
}

func TestNewMediaServerRequestJoinsPathCorrectly(t *testing.T) {
	baseURL, path := "http://127.0.0.1/api/plex/", "v2?a=b"
	client, _ := New("test")
	client.MediaServerURL = baseURL

	req, err := client.NewMediaServerRequest(http.MethodGet, path, nil)

	if assert.NoError(t, err, "NewMediaServerRequest should not return an error") {
		assert.Equal(t, fmt.Sprintf("%s%s", baseURL, path), req.URL.String(), "NewMediaServerRequest should join the base URL and path")
	}
}
