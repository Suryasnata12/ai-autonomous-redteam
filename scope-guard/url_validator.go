package scopeguard

import (
	"errors"
	"net"
	"net/url"
	"strings"
)

func ValidateURL(raw string) (*url.URL, error) {
	u, err := url.Parse(raw)
	if err != nil {
		return nil, err
	}

	if u.Scheme != "http" && u.Scheme != "https" {
		return nil, errors.New("invalid scheme")
	}

	if u.User != nil {
		return nil, errors.New("userinfo not allowed")
	}

	if u.Hostname() == "" {
		return nil, errors.New("empty hostname")
	}

	// Block raw IPs (initial MVP)
	if net.ParseIP(u.Hostname()) != nil {
		return nil, errors.New("IP literals not allowed")
	}

	if strings.Contains(u.Hostname(), "..") {
		return nil, errors.New("invalid hostname")
	}

	return u, nil
}
