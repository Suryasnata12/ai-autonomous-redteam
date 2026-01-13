package orchestrator

import "ai-autonomous-redteam/recon-engine/model"

// Every recon module MUST implement this
type ReconModule interface {
	Name() string
	Run(domain string, result *model.ReconResult) error
}

// Orchestrates recon modules safely
func RunRecon(domain string, modules []ReconModule) (*model.ReconResult, error) {
	result := &model.ReconResult{
		Domain:  domain,
		Headers: make(map[string]string),
	}

	for _, m := range modules {
		if err := m.Run(domain, result); err != nil {
			result.Notes = append(
				result.Notes,
				m.Name()+": "+err.Error(),
			)
		}
	}

	return result, nil
}
