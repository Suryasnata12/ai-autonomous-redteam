package orchestrator

import (
	"ai-autonomous-redteam/recon-engine/modules"
)

// DefaultModules returns the standard recon pipeline
func DefaultModules() []ReconModule {
	return []ReconModule{
		&modules.DNSModule{},
		&modules.HTTPProbe{},
		&modules.HeadersModule{},
	}
}
