package scopeguard

import (
	"net/http"
	"time"
)

func NewSafeClient() *http.Client {
	return &http.Client{
		Timeout: 10 * time.Second,

		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// Re-validate redirects
			_, err := ValidateURL(req.URL.String())
			if err != nil {
				return err
			}

			_, err = ResolveAndCheck(req.URL.Hostname())
			return err
		},
	}
}
