package safety

import "ai-autonomous-redteam/safety/model"

type BudgetManager struct {
	Budget model.Budget
	State  model.BudgetState
}

func NewBudgetManager(b model.Budget) *BudgetManager {
	return &BudgetManager{
		Budget: b,
		State: model.BudgetState{
			UsedByAttack: make(map[string]int),
		},
	}
}

func (m *BudgetManager) Allow(attack string) bool {
	if m.State.TotalUsed >= m.Budget.MaxTotalRequests {
		return false
	}

	if m.State.UsedByAttack[attack] >= m.Budget.MaxPerAttack[attack] {
		return false
	}

	m.State.TotalUsed++
	m.State.UsedByAttack[attack]++
	return true
}
