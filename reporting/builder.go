package reporting

import (
	plan "ai-autonomous-redteam/ai-planner/model"
	reason "ai-autonomous-redteam/ai-reasoner/model"
	recon "ai-autonomous-redteam/recon-engine/model"
	"ai-autonomous-redteam/reporting/model"
)

func BuildReport(
	target string,
	r recon.ReconResult,
	re reason.ReasoningResult,
	p plan.ExploitPlan,
) model.Report {

	decision := "STOP"
	reasonText := "No safe exploit paths identified"

	if len(p.Steps) > 0 {
		decision = "PROCEED (PLANNED ONLY)"
		reasonText = "Safe exploit plan generated within budget"
	}

	return model.Report{
		Target:         target,
		Recon:          r,
		Reasoning:      re,
		Plan:           p,
		FinalDecision:  decision,
		DecisionReason: reasonText,
	}
}
