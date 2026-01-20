package main

import (
	"fmt"
	"time"

	"ai-autonomous-redteam/learning/model"
	"ai-autonomous-redteam/learning/store"
)

func main() {
	run := model.RunRecord{
		ID:         "test-run-1",
		Target:     "example.com",
		Timestamp:  time.Now(),
		Outcome:    model.OutcomeBlocked,
		StopReason: "WAF triggered",
	}

	_ = store.SaveRun(run)

	history, _ := store.LoadAllRuns()

	fmt.Printf("Loaded %d historical runs\n", len(history))
}
