package modules

import "ai-autonomous-redteam/recon-engine/model"

type AuthDetector struct{}

func (a *AuthDetector) Name() string {
	return "auth-detector"
}

func (a *AuthDetector) Run(domain string, result *model.ReconResult) error {
	headers := result.Headers

	auth := model.AuthInfo{}

	// ---- Cookie-based auth ----
	if headers["Set-Cookie"] != "" {
		auth.UsesCookies = true
		auth.AuthHints = append(auth.AuthHints, "Set-Cookie present")
	}

	// ---- JWT / Bearer tokens ----
	if headers["Authorization"] != "" {
		auth.UsesJWT = true
		auth.AuthHints = append(auth.AuthHints, "Authorization header observed")
	}

	// ---- CSRF indicators ----
	if headers["X-CSRF-Token"] != "" ||
		headers["X-XSRF-Token"] != "" {
		auth.HasCSRF = true
		auth.AuthHints = append(auth.AuthHints, "CSRF header detected")
	}

	// ---- Framework hints ----
	if contains(result.TechStack, "Next.js") {
		auth.AuthHints = append(auth.AuthHints, "Likely NextAuth / cookie session")
	}

	result.Auth = auth
	return nil
}

// helper
func contains(list []string, v string) bool {
	for _, x := range list {
		if x == v {
			return true
		}
	}
	return false
}
