package model

type ReasoningResult struct {
	TargetType     string
	RiskLevel      string
	AllowedAttacks []AttackHypothesis
	AvoidAttacks   []string
	Notes          []string
}

type AttackHypothesis struct {
	Class      string
	Confidence float64
	Rationale  string
}
