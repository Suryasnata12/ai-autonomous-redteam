package ownership

import "time"

type VerifiedDomain struct {
	Domain    string
	VerifiedAt time.Time
	ExpiresAt time.Time
}

var store = make(map[string]VerifiedDomain)

func SaveVerified(domain string) {
	store[domain] = VerifiedDomain{
		Domain: domain,
		VerifiedAt: time.Now(),
		ExpiresAt: time.Now().Add(30 * 24 * time.Hour),
	}
}

func IsVerified(domain string) bool {
	v, ok := store[domain]
	return ok && time.Now().Before(v.ExpiresAt)
}
