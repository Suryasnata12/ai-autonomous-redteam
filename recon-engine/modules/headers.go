package modules

import "ai-autonomous-redteam/recon-engine/model"

type HeadersModule struct{}

func (h *HeadersModule) Name() string {
	return "headers"
}

func (h *HeadersModule) Run(domain string, result *model.ReconResult) error {
	if server, ok := result.Headers["Server"]; ok {
		result.Server = server
	}
	return nil
}
