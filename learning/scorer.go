package learning

import "ai-autonomous-redteam/learning/model"

func ScoreAttackClass(class string, history []model.RunRecord) float64 {
	score := 1.0

	for _, run := range history {
		for _, step := range run.Plan.Steps {
			if step.AttackClass == class {
				if run.Outcome == model.OutcomeBlocked {
					score -= 0.3
				}
				if run.Outcome == model.OutcomeStopped {
					score -= 0.1
				}
			}
		}
	}

	if score < 0.1 {
		score = 0.1
	}

	return score
}
