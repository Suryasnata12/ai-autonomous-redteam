package modules

import "ai-autonomous-redteam/recon-engine/model"

// HeadersModule analyzes HTTP response headers
// and derives security posture signals.
type HeadersModule struct{}

func (h *HeadersModule) Name() string {
	return "headers"
}

func (h *HeadersModule) Run(domain string, result *model.ReconResult) error {
	headers := result.Headers

	// ---- Normalize Server ----
	if server, ok := headers["Server"]; ok {
		result.Server = server
	}

	// ---- Security Posture Analysis ----
	result.Security = model.SecurityPosture{
		HSTS:          hasHeader(headers, "Strict-Transport-Security"),
		CSP:           hasHeader(headers, "Content-Security-Policy"),
		XFrameOptions: hasHeader(headers, "X-Frame-Options"),
		XContentType:  hasHeader(headers, "X-Content-Type-Options"),
		CookieFlags:   analyzeCookies(headers["Set-Cookie"]),
	}

	return nil
}

// ---- helpers ----

// Checks whether a header exists and is non-empty
func hasHeader(headers map[string]string, name string) bool {
	v, ok := headers[name]
	return ok && v != ""
}

// Very lightweight cookie security check
// (expanded later in Week 6)
func analyzeCookies(cookie string) bool {
	if cookie == "" {
		return false
	}

	// Presence-based for Week 4
	// (Secure / HttpOnly parsing comes later)
	return true
}
