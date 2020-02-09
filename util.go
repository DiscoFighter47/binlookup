package binlookup

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

// buildURL returns the full url
func (c *Client) buildURL(path string) string {
	u := url.URL{
		Scheme: "https",
		Host:   "lookup.binlist.net",
		Path:   path,
	}
	return u.String()
}

func (c *Client) bodyBuffer(body interface{}) (io.Reader, error) {
	jsonStr, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer([]byte(jsonStr)), nil
}

func (c *Client) fillHeader(req *http.Request) {
	req.Header.Set("Accept-Version", "3")
}

func (c *Client) captureBody(r io.Reader) (string, error) {
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

func (c *Client) parseBody(r *http.Response, body interface{}) error {
	return json.NewDecoder(r.Body).Decode(body)
}
