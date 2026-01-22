package learning

import (
	reasonModel "ai-autonomous-redteam/ai-reasoner/model"
	lmodel "ai-autonomous-redteam/learning/model"
	"ai-autonomous-redteam/learning/store"
)

// ApplyLearning adjusts reasoning using historical run outcomes.
// Conservative rule: repeated STOPPED outcomes reduce risk appetite.
func ApplyLearning(reason reasonModel.ReasoningResult) reasonModel.ReasoningResult {
	runs, err := store.LoadAllRuns()
	if err != nil || len(runs) == 0 {
		return reason
	}

	stoppedCount := 0

	for _, run := range runs {
		if run.Outcome == lmodel.OutcomeStopped {
			stoppedCount++
		}
	}

	// If system frequently stops, downgrade risk
	if stoppedCount >= 3 {
		reason.RiskLevel = "low"
		reason.Notes = append(
			reason.Notes,
			"Learning: repeated STOPPED outcomes in historical runs",
		)
	}

	return reason
}
