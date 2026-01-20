package main

import (
	"fmt"

	"ai-autonomous-redteam/ownership"
	"ai-autonomous-redteam/recon-engine/modules"
	"ai-autonomous-redteam/recon-engine/orchestrator"
)

func main() {
	domain := "nextjs.org"
	ownership.SaveVerified(domain)

	// ---- Enforcement check ----
	if !ownership.IsVerified(domain) {
		panic("❌ domain not verified – recon blocked")
	}

	// ---- Recon modules ----
	modulesList := []orchestrator.ReconModule{
		&modules.DNSModule{},
		&modules.HTTPProbe{},
		&modules.HeadersModule{},
		&modules.HTTPFingerprint{},
		&modules.WAFDetector{},
		&modules.JSCrawler{},
		&modules.AuthDetector{},
	}

	// ---- Run recon ----
	result, err := orchestrator.RunRecon(domain, modulesList)
	if err != nil {
		panic(err)
	}

	fmt.Printf("\n=== RECON RESULT ===\n%+v\n", result)
}
