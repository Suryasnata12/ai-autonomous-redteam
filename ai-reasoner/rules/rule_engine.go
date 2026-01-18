package rules

import (
	"ai-autonomous-redteam/ai-reasoner/model"
	recon "ai-autonomous-redteam/recon-engine/model"
)

func ApplyRules(r *recon.ReconResult) model.ReasoningResult {
	result := model.ReasoningResult{
		TargetType: "unknown",
		RiskLevel:  "low",
	}

	// ---- Target classification ----
	if contains(r.TechStack, "Next.js") {
		result.TargetType = "modern web app"
	}

	// ---- Security posture ----
	if r.Security.CSP && r.Security.HSTS {
		result.Notes = append(result.Notes, "Strong transport and CSP detected")
		result.AvoidAttacks = append(result.AvoidAttacks, "Reflected XSS")
	}

	// ---- WAF influence ----
	if r.WAF != "unknown" {
		result.Notes = append(result.Notes, "WAF detected: "+r.WAF)
		result.AvoidAttacks = append(result.AvoidAttacks, "Brute force")
	}

	// ---- Auth reasoning ----
	if r.Auth.UsesCookies {
		result.AllowedAttacks = append(result.AllowedAttacks,
			model.AttackHypothesis{
				Class:      "CSRF",
				Confidence: 0.7,
				Rationale:  "Cookie-based auth observed",
			},
		)
	}

	if len(r.Endpoints) > 0 {
		result.AllowedAttacks = append(result.AllowedAttacks,
			model.AttackHypothesis{
				Class:      "API misuse",
				Confidence: 0.6,
				Rationale:  "Client-discovered API endpoints",
			},
		)
	}

	return result
}

func contains(list []string, v string) bool {
	for _, x := range list {
		if x == v {
			return true
		}
	}
	return false
}
