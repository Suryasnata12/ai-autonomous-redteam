package modules

import (
	"ai-autonomous-redteam/recon-engine/model"
	scopeguard "ai-autonomous-redteam/scope-guard"
)

type HTTPProbe struct{}

func (h *HTTPProbe) Name() string {
	return "http-probe"
}

func (h *HTTPProbe) Run(domain string, result *model.ReconResult) error {
	client := scopeguard.NewSafeClient()

	resp, err := client.Head("https://" + domain)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	for k, v := range resp.Header {
		if len(v) > 0 {
			result.Headers[k] = v[0]
		}
	}

	result.Server = resp.Header.Get("Server")
	return nil
}
