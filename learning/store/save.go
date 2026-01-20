package store

import (
	"encoding/json"
	"os"
	"path/filepath"

	"ai-autonomous-redteam/config"
	"ai-autonomous-redteam/learning/model"
)

func SaveRun(run model.RunRecord) error {
	dir := config.DataDir()
	_ = os.MkdirAll(dir, 0755)

	path := filepath.Join(dir, run.ID+".json")

	data, err := json.MarshalIndent(run, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
