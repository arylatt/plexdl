package plexapi

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTokenAuthorizer(t *testing.T) {
	token := "test"

	tokenAuthorizer := NewTokenAuthorizer(token)

	assert.Equal(t, token, tokenAuthorizer.token, "NewTokenAuthorizer should return the same token it was passed in")
	assert.Equal(t, http.DefaultTransport, tokenAuthorizer.RoundTripper, "NewTokenAuthorizer should return the http.DefaultTransport")
}

type MockBaseTransport struct {
	test  *testing.T
	token string
}

func (m *MockBaseTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	reqQuery := req.URL.Query()
	assert.Equal(m.test, m.token, reqQuery.Get("X-Plex-Token"), "X-Plex-Token on request passed to parent RoundTripper should have value from TokenAuthorizer")

	return nil, nil
}

func TestTokenAuthorizerRoundTrip(t *testing.T) {
	token := "test"

	tokenAuthorizer := &TokenAuthorizer{
		RoundTripper: &MockBaseTransport{test: t, token: token},
		token:        token,
	}

	mockReq, err := http.NewRequest("GET", "http://127.0.0.1", nil)
	assert.NoError(t, err, "http.NewRequest should not return an error")

	tokenAuthorizer.RoundTrip(mockReq)
}

func TestTokenAuthorizerToken(t *testing.T) {
	token := "test"

	tokenAuthorizer := NewTokenAuthorizer(token)

	assert.Equal(t, token, tokenAuthorizer.Token(), "TokenAuthorizer.Token should return the same token as was passed in")
}
