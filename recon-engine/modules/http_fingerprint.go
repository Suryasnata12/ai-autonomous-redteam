package modules

import "ai-autonomous-redteam/recon-engine/model"

type HTTPFingerprint struct{}

func (h *HTTPFingerprint) Name() string {
	return "http-fingerprint"
}

func (h *HTTPFingerprint) Run(domain string, result *model.ReconResult) error {
	headers := result.Headers

	// Framework hints
	if v, ok := headers["X-Powered-By"]; ok {
		result.TechStack = append(result.TechStack, v)
	}

	if _, ok := headers["X-Nextjs-Cache"]; ok {
		result.TechStack = append(result.TechStack, "Next.js")
	}

	if _, ok := headers["X-React-Server"]; ok {
		result.TechStack = append(result.TechStack, "React")
	}

	if result.Server != "" {
		result.TechStack = append(result.TechStack, "Server:"+result.Server)
	}

	return nil
}
