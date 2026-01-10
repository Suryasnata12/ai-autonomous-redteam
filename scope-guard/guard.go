package scopeguard

import (
	"net"
)

func ResolveAndCheck(host string) ([]net.IP, error) {
	ips, err := net.LookupIP(host)
	if err != nil {
		return nil, err
	}

	for _, ip := range ips {
		if IsBlockedIP(ip) {
			return nil, ErrBlockedIP
		}
	}
	return ips, nil
}
