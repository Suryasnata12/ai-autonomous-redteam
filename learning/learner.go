package learning

import (
	reason "ai-autonomous-redteam/ai-reasoner/model"
	"ai-autonomous-redteam/learning/model"
)

func AdjustReasoning(
	r reason.ReasoningResult,
	history []model.RunRecord,
) reason.ReasoningResult {

	for i, attack := range r.AllowedAttacks {
		factor := ScoreAttackClass(attack.Class, history)
		r.AllowedAttacks[i].Confidence *= factor
	}

	return r
}
