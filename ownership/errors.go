package ownership

import "errors"

var (
	// Returned when neither DNS nor HTTP verification succeeds
	ErrVerificationFailed = errors.New("ownership verification failed")

	// Returned when verification token has expired
	ErrTokenExpired = errors.New("verification token expired")

	// Returned when domain is not verified or verification expired
	ErrDomainNotVerified = errors.New("domain ownership not verified")
)
