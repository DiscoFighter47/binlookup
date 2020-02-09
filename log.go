package binlookup

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/DiscoFighter47/zlog"
)

func (c *Client) logRequest(r *http.Request) {
	zlog.Infof("[binlist] Req: %s %v", r.Method, r.URL)
}

func (c *Client) logResponse(r *http.Response, t time.Duration) {
	zlog.Infof("[binlist] Res: %d %s", r.StatusCode, t.String())
	if r.Body != nil {
		var buf bytes.Buffer
		tee := io.TeeReader(r.Body, &buf)
		if body, err := c.captureBody(tee); err == nil {
			if len(body) == 0 {
				body = "[]"
			}
			zlog.Infof("[binlist] Res Body: %s", body)
		}
		r.Body = ioutil.NopCloser(&buf)
	}
}
