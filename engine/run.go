package engine

import (
	scopeguard "ai-autonomous-redteam/scope-guard"

	aiplanner "ai-autonomous-redteam/ai-planner"
	aireasoner "ai-autonomous-redteam/ai-reasoner"
	reconEngine "ai-autonomous-redteam/recon-engine/orchestrator"

	"ai-autonomous-redteam/learning"

	"ai-autonomous-redteam/safety"
	smodel "ai-autonomous-redteam/safety/model"

	"ai-autonomous-redteam/reporting"
	rmodel "ai-autonomous-redteam/reporting/model"
)

func Run(target string) (rmodel.Report, error) {

	// 1. Scope guard
	if _, err := scopeguard.ResolveAndCheck(target); err != nil {
		return rmodel.Report{}, err
	}

	// 2. Ownership
	/*if !ownership.IsVerified(target) {
		return rmodel.Report{}, errors.New("ownership not verified")
	}*/

	// 3. Recon
	recon := reconEngine.Execute(target)

	// 4. Reasoning
	reason := aireasoner.Reason(&recon)

	// 5. Learning (RESTORED, SAFE)
	reason = learning.ApplyLearning(reason)

	// 6. Planning
	plan := aiplanner.Plan(target, reason)

	// 7. Safety gate
	budget := smodel.Budget{
		MaxTotalRequests: 5,
		MaxPerAttack: map[string]int{
			"XSS":  1,
			"CSRF": 1,
		},
	}

	if !safety.AllowPlan(plan, budget) {
		plan.Steps = nil
	}

	// 8. Build report
	return reporting.BuildReport(
		target,
		recon,
		reason,
		plan,
	), nil
}
