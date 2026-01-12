package ownership

import (
	"fmt"
	"net"
	"strings"
)

func VerifyDNS(domain, token string) error {
	txts, err := net.LookupTXT("_redforge." + domain)
	if err != nil {
		return err
	}

	expected := fmt.Sprintf("redforge-verification=%s", token)

	for _, t := range txts {
		if strings.TrimSpace(t) == expected {
			return nil
		}
	}
	return ErrVerificationFailed
}
