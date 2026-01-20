package main

import (
	"fmt"

	aireasoner "ai-autonomous-redteam/ai-reasoner"
	recon "ai-autonomous-redteam/recon-engine/model"
)

func main() {
	// ---- Mock recon result (from Weeks 1–6) ----
	reconResult := &recon.ReconResult{
		TechStack: []string{"Next.js", "Server:Vercel"},
		WAF:       "Cloudflare",
		Endpoints: []string{},
		Security: recon.SecurityPosture{
			HSTS: true,
			CSP:  true,
		},
		Auth: recon.AuthInfo{
			UsesCookies: false,
			UsesJWT:     false,
			HasCSRF:     false,
		},
	}

	// ---- Run AI Reasoner ----
	result := aireasoner.Reason(reconResult)

	fmt.Printf("\n=== AI REASONING ===\n%+v\n", result)
}
