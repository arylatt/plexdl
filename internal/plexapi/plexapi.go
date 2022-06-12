package plexapi

import (
	"context"
	"encoding/xml"
	"io"
	"net/http"
	"net/url"
)

// Client is an API client for interacting with Plex APIs.
type Client struct {
	// MediaServerURL is the HOST:PORT of a Plex Media Server to issue requests to
	MediaServerURL string

	ctx    context.Context
	client *http.Client
}

// NewWithContext returns a new Client with a supplied context and optional token.
func NewWithContext(ctx context.Context, token ...string) (client *Client, err error) {
	httpClient := &http.Client{}

	if len(token) >= 1 {
		httpClient.Transport = NewTokenAuthorizer(token[0])
	}

	client = &Client{
		ctx:    ctx,
		client: httpClient,
	}

	return
}

// New returns a new Client with the background context and an optional token.
func New(token ...string) (*Client, error) {
	return NewWithContext(context.Background(), token...)
}

// DecodeXMLResponse is a helper function to decode XML to a struct
func DecodeXMLResponse(resp *http.Response, v any) error {
	defer resp.Body.Close()
	return xml.NewDecoder(resp.Body).Decode(v)
}

// Client.SetTransport allows overriding the HTTP client Transport with a custom http.RoundTripper.
// Can be used to pass in a TokenAuthorizer or a custom http.RoundTripper interface.
func (c *Client) SetTransport(rt http.RoundTripper) {
	c.client.Transport = rt
}

// Client.NewMediaServerRequestWithContext returns a new http.Request using the Client.MediaServerURL as a base using a custom request context.
func (c *Client) NewMediaServerRequestWithContext(ctx context.Context, method, path string, body io.Reader) (req *http.Request, err error) {
	baseURL, err := url.Parse(c.MediaServerURL)
	if err != nil {
		return
	}

	reqURL, err := baseURL.Parse(path)
	if err != nil {
		return
	}

	req, err = http.NewRequestWithContext(ctx, method, reqURL.String(), body)
	return
}

// Client.NewMediaServerRequest returns a new http.Request using the Client.MediaServerURL as a base.
func (c *Client) NewMediaServerRequest(method, path string, body io.Reader) (*http.Request, error) {
	return c.NewMediaServerRequestWithContext(c.ctx, method, path, body)
}

func (c *Client) MediaServerGet(path string) (resp *http.Response, err error) {
	req, err := c.NewMediaServerRequest(http.MethodGet, path, nil)
	if err != nil {
		return
	}

	return c.Do(req)
}

// Client.Do sends a http.Request and returns the response and/or error from the underlying HTTP client.
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	return c.client.Do(req)
}
