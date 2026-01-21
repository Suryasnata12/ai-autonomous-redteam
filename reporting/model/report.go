package model

import (
	plan "ai-autonomous-redteam/ai-planner/model"
	reason "ai-autonomous-redteam/ai-reasoner/model"
	recon "ai-autonomous-redteam/recon-engine/model"
)

type Report struct {
	Target string

	Recon     recon.ReconResult
	Reasoning reason.ReasoningResult
	Plan      plan.ExploitPlan

	FinalDecision  string
	DecisionReason string
}
