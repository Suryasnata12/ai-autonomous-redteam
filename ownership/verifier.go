package ownership

func Verify(domain string, token string) error {
	if err := VerifyDNS(domain, token); err == nil {
		SaveVerified(domain)
		return nil
	}

	if err := VerifyHTTP(domain, token); err == nil {
		SaveVerified(domain)
		return nil
	}

	return ErrVerificationFailed
}
