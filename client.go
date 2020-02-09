package binlookup

import (
	"net/http"
	"time"

	"github.com/gojektech/heimdall/httpclient"
)

// Client represents host and access token for accessing wallet api
type Client struct {
	client *httpclient.Client
}

// NewClient returns a new client instance
func NewClient(t time.Duration) *Client {
	c := &Client{
		client: httpclient.NewClient(httpclient.WithHTTPTimeout(t)),
	}
	return c
}

// Do actually makes the http request
func (c *Client) do(url string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, c.buildURL(url), nil)
	if err != nil {
		return nil, err
	}
	c.fillHeader(req)
	c.logRequest(req)
	t := time.Now()
	resp, err := c.client.Do(req)
	if err == nil {
		c.logResponse(resp, time.Since(t))
	}
	return resp, err
}
