package plexapi

import "net/http"

const PlexTokenAuthBaseURL = "https://plex.tv/api/v2/"

// TokenAuthorizer is an implementation of the http.RoundTripper interface.
// It will attach an "X-Plex-Token" query parameter to outgoing requests
// with the value of token supplied when calling NewTokenAuthorizer.
type TokenAuthorizer struct {
	http.RoundTripper
	token string
}

// NewTokenAuthorizer creates a new TokenAuthorizer with the supplied token.
func NewTokenAuthorizer(token string) *TokenAuthorizer {
	return &TokenAuthorizer{
		RoundTripper: http.DefaultTransport,
		token:        token,
	}
}

// TokenAuthorizer.RoundTrip will attach the TokenAuthorizer token to a given
// HTTP request, and then call http.DefaultTransport.RoundTrip.
func (ta *TokenAuthorizer) RoundTrip(req *http.Request) (*http.Response, error) {
	reqQuery := req.URL.Query()

	reqQuery.Add("X-Plex-Token", ta.token)

	req.URL.RawQuery = reqQuery.Encode()

	return ta.RoundTripper.RoundTrip(req)
}

// TokenAuthorizer.Token will return the current token.
func (ta *TokenAuthorizer) Token() string {
	return ta.token
}
