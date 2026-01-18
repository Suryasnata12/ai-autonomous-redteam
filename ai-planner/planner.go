package aiplanner

import (
	"ai-autonomous-redteam/ai-planner/model"
	reason "ai-autonomous-redteam/ai-reasoner/model"
)

func Plan(target string, r reason.ReasoningResult) model.ExploitPlan {
	plan := model.ExploitPlan{
		Target:      target,
		MaxRequests: 5, // default conservative budget
	}

	// ---- Stop conditions (global safety) ----
	plan.StopConditions = append(plan.StopConditions,
		"WAF detected",
		"Auth lockout",
		"Unexpected 4xx/5xx spike",
	)

	// ---- Translate allowed attacks into steps ----
	for _, attack := range r.AllowedAttacks {
		switch attack.Class {

		case "CSRF":
			plan.Steps = append(plan.Steps, model.PlanStep{
				AttackClass: "CSRF",
				Description: "Attempt CSRF on state-changing endpoints",
				MaxAttempts: 2,
			})

		case "API misuse":
			plan.Steps = append(plan.Steps, model.PlanStep{
				AttackClass: "API misuse",
				Description: "Test API endpoints for missing auth or IDOR",
				MaxAttempts: 3,
			})
		}
	}

	return plan
}
