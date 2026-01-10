package scopeguard

import (
	"net/http"
	"time"
)

// ---- helper (LOCAL to this file) ----

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r)
}

// ---- safe HTTP client ----

func NewSafeClient() *http.Client {
	return &http.Client{
		Timeout: 10 * time.Second,

		Transport: roundTripperFunc(func(req *http.Request) (*http.Response, error) {
			host := req.URL.Hostname()

			if !AllowRequest(host) {
				return nil, ErrRateLimit
			}

			return http.DefaultTransport.RoundTrip(req)
		}),

		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if _, err := ValidateURL(req.URL.String()); err != nil {
				return err
			}

			_, err := ResolveAndCheck(req.URL.Hostname())
			return err
		},
	}
}
