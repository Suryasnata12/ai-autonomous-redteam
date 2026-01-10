package scopeguard

import "net"

var blockedRanges []*net.IPNet

func init() {
	ranges := []string{
		"127.0.0.0/8",
		"10.0.0.0/8",
		"172.16.0.0/12",
		"192.168.0.0/16",
		"169.254.0.0/16", // link-local
		"::1/128",
		"fc00::/7",
		"fe80::/10",
	}

	for _, cidr := range ranges {
		_, block, _ := net.ParseCIDR(cidr)
		blockedRanges = append(blockedRanges, block)
	}
}

func IsBlockedIP(ip net.IP) bool {
	for _, block := range blockedRanges {
		if block.Contains(ip) {
			return true
		}
	}
	return false
}
