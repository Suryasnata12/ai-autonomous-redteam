package orchestrator

import "ai-autonomous-redteam/recon-engine/model"

// Execute is a high-level recon entry point used by the engine.
// It runs recon using the default module pipeline.
func Execute(domain string) model.ReconResult {
	modules := DefaultModules()

	result, err := RunRecon(domain, modules)
	if err != nil {
		return model.ReconResult{
			Domain: domain,
			Notes:  []string{"recon failed: " + err.Error()},
		}
	}

	return *result
}
