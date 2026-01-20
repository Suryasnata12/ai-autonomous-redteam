package model

type Budget struct {
	MaxTotalRequests int
	MaxPerAttack     map[string]int
}

type BudgetState struct {
	TotalUsed    int
	UsedByAttack map[string]int
}
