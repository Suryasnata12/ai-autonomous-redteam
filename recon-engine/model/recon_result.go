package model

type ReconResult struct {
	Domain string
	IPs    []string

	Headers map[string]string
	Server  string

	TechStack []string
	WAF       string
	Security  SecurityPosture

	Notes     []string
	JSFiles   []string
	Endpoints []string
}

type SecurityPosture struct {
	HSTS          bool
	CSP           bool
	XFrameOptions bool
	XContentType  bool
	CookieFlags   bool
}
