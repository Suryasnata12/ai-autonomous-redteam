package config

import (
	"os"
	"path/filepath"
)

// Assumes execution anywhere inside repo
func ProjectRoot() string {
	dir, _ := os.Getwd()

	for {
		if filepath.Base(dir) == "ai-autonomous-redteam" {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}

	// fallback (safe)
	return "."
}
