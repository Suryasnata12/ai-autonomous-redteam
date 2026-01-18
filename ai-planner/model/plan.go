package model

type ExploitPlan struct {
	Target         string
	MaxRequests    int
	Steps          []PlanStep
	StopConditions []string
}

type PlanStep struct {
	AttackClass string
	Description string
	MaxAttempts int
}
