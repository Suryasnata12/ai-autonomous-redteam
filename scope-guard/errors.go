package scopeguard

import "errors"

var ErrBlockedIP = errors.New("blocked IP range")
var ErrRateLimit = errors.New("rate limit exceeded")
