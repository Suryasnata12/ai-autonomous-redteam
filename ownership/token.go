package ownership

import (
	"crypto/rand"
	"encoding/hex"
	"time"
)

type VerificationToken struct {
	Domain    string
	Token     string
	ExpiresAt time.Time
}

func GenerateToken(domain string) VerificationToken {
	b := make([]byte, 16)
	rand.Read(b)

	return VerificationToken{
		Domain:    domain,
		Token:     hex.EncodeToString(b),
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}
}
