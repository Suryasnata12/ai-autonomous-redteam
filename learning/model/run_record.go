package model

import (
	"time"

	plan "ai-autonomous-redteam/ai-planner/model"
	reason "ai-autonomous-redteam/ai-reasoner/model"
	recon "ai-autonomous-redteam/recon-engine/model"
)

type RunOutcome string

const (
	OutcomeStopped RunOutcome = "STOPPED"
	OutcomeNoOp    RunOutcome = "NO_OP"
	OutcomeBlocked RunOutcome = "BLOCKED"
)

type RunRecord struct {
	ID        string
	Target    string
	Timestamp time.Time

	Recon     recon.ReconResult
	Reasoning reason.ReasoningResult
	Plan      plan.ExploitPlan

	Outcome    RunOutcome
	StopReason string
}
