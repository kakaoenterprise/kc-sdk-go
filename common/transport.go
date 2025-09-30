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

func newAuthedHTTPClient(base *http.Client, token, ua string) (*http.Client, *xAuthTokenTransport) {
	if base == nil {
		base = http.DefaultClient
	}
	rt := base.Transport
	if rt == nil {
		rt = http.DefaultTransport
	}
	auth := &xAuthTokenTransport{base: rt}
	auth.token.Store(token)
	auth.userAgent.Store(ua)

	cli := &http.Client{
		Transport: auth,
		Timeout:   base.Timeout,
	}
	return cli, auth
}
