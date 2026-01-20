package main

import (
	"fmt"

	aiplanner "ai-autonomous-redteam/ai-planner"
	reason "ai-autonomous-redteam/ai-reasoner/model"
)

func main() {
	reasoning := reason.ReasoningResult{
		AllowedAttacks: []reason.AttackHypothesis{
			{
				Class:      "CSRF",
				Confidence: 0.7,
			},
		},
	}

	plan := aiplanner.Plan("example.com", reasoning)

	fmt.Printf("\n=== EXPLOIT PLAN ===\n%+v\n", plan)
}
