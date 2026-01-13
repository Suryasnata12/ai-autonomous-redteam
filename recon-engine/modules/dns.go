package modules

import (
	"net"

	"ai-autonomous-redteam/recon-engine/model"
)

type DNSModule struct{}

func (d *DNSModule) Name() string {
	return "dns"
}

func (d *DNSModule) Run(domain string, result *model.ReconResult) error {
	ips, err := net.LookupIP(domain)
	if err != nil {
		return err
	}

	for _, ip := range ips {
		result.IPs = append(result.IPs, ip.String())
	}
	return nil
}
