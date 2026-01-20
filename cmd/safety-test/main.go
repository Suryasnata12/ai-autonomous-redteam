package main

import (
	"fmt"

	"ai-autonomous-redteam/ai-planner/model"
	"ai-autonomous-redteam/safety"
	smodel "ai-autonomous-redteam/safety/model"
)

func main() {
	plan := model.ExploitPlan{
		MaxRequests: 5,
		Steps: []model.PlanStep{
			{AttackClass: "CSRF", MaxAttempts: 2},
			{AttackClass: "API misuse", MaxAttempts: 4},
		},
	}

	budget := smodel.Budget{
		MaxTotalRequests: 5,
		MaxPerAttack: map[string]int{
			"CSRF":       2,
			"API misuse": 3,
		},
	}

	allowed := safety.AllowPlan(plan, budget)
	fmt.Println("Plan allowed:", allowed)
}
