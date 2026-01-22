package model

import (
	plan "ai-autonomous-redteam/ai-planner/model"
	reason "ai-autonomous-redteam/ai-reasoner/model"
	recon "ai-autonomous-redteam/recon-engine/model"
)

// Report is the final consolidated, read-only output
// exposed to the UI and API.
type Report struct {
	Target string `json:"target"`

	Recon     recon.ReconResult      `json:"recon"`
	Reasoning reason.ReasoningResult `json:"reasoning"`
	Plan      plan.ExploitPlan       `json:"plan"`

	FinalDecision  string `json:"final_decision"`
	DecisionReason string `json:"decision_reason"`
}
