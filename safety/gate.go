package safety

import (
	"ai-autonomous-redteam/ai-planner/model"
	sm "ai-autonomous-redteam/safety/model"
)

func AllowPlan(plan model.ExploitPlan, budget sm.Budget) bool {
	total := 0
	for _, step := range plan.Steps {
		total += step.MaxAttempts
	}

	return total <= budget.MaxTotalRequests
}
