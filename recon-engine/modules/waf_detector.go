package modules

import "ai-autonomous-redteam/recon-engine/model"

type WAFDetector struct{}

func (w *WAFDetector) Name() string {
	return "waf-detector"
}

func (w *WAFDetector) Run(domain string, result *model.ReconResult) error {
	headers := result.Headers

	switch {
	case headers["Server"] == "cloudflare":
		result.WAF = "Cloudflare"

	case headers["X-Akamai-Transformed"] != "":
		result.WAF = "Akamai"

	case headers["X-CDN"] != "":
		result.WAF = headers["X-CDN"]

	default:
		result.WAF = "unknown"
	}

	return nil
}
