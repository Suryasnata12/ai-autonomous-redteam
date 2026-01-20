package config

import "path/filepath"

func DataDir() string {
	return filepath.Join(ProjectRoot(), "learning", "runs")
}
