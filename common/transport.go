package common

import (
	"net/http"
	"sync/atomic"
)

type xAuthTokenTransport struct {
	base      http.RoundTripper
	token     atomic.Value
	userAgent atomic.Value
}

func (t *xAuthTokenTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	r := req.Clone(req.Context())
	if tok, ok := t.token.Load().(string); ok && tok != "" {
		r.Header.Set("X-Auth-Token", tok)
	}
	if ua, ok := t.userAgent.Load().(string); ok && ua != "" {
		r.Header.Set("User-Agent", ua)
	}
	return t.base.RoundTrip(r)
}
