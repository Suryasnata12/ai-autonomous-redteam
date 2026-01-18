package aireasoner

import (
	"ai-autonomous-redteam/ai-reasoner/model"
	"ai-autonomous-redteam/ai-reasoner/rules"
	recon "ai-autonomous-redteam/recon-engine/model"
)

func Reason(r *recon.ReconResult) model.ReasoningResult {
	return rules.ApplyRules(r)
}
