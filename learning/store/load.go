package store

import (
	"encoding/json"
	"os"
	"path/filepath"

	"ai-autonomous-redteam/learning/model"
)

func LoadAllRuns() ([]model.RunRecord, error) {
	var runs []model.RunRecord

	files, err := os.ReadDir("learning/runs")
	if err != nil {
		return runs, nil // no history yet
	}

	for _, f := range files {
		data, err := os.ReadFile(filepath.Join("learning/runs", f.Name()))
		if err != nil {
			continue
		}

		var run model.RunRecord
		if err := json.Unmarshal(data, &run); err == nil {
			runs = append(runs, run)
		}
	}

	return runs, nil
}
