package main

import (
	"time"

	"ai-autonomous-redteam/reporting"
	rmodel "ai-autonomous-redteam/reporting/model"

	plan "ai-autonomous-redteam/ai-planner/model"
	reason "ai-autonomous-redteam/ai-reasoner/model"
	recon "ai-autonomous-redteam/recon-engine/model"
)

func main() {
	report := rmodel.Report{
		Target: "example.com",

		Recon: recon.ReconResult{
			TechStack: []string{"Next.js", "Server:Vercel"},
			WAF:       "Cloudflare",
		},

		Reasoning: reason.ReasoningResult{
			TargetType: "modern web app",
			RiskLevel:  "low",
			AvoidAttacks: []string{
				"Reflected XSS",
				"Brute force",
			},
			Notes: []string{
				"Strong CSP detected",
				"WAF detected: Cloudflare",
			},
		},

		Plan: plan.ExploitPlan{
			Target:      "example.com",
			MaxRequests: 5,
			Steps:       []plan.PlanStep{}, // empty = STOP
		},

		FinalDecision:  "STOP",
		DecisionReason: "No safe exploit paths identified",
	}

	reporting.StartServer(report)

	// Keep server alive
	time.Sleep(time.Hour)
}
