package scopeguard

import (
	"sync"
	"time"
)

type limiter struct {
	count     int
	windowEnd time.Time
}

var (
	mu       sync.Mutex
	limiters = make(map[string]*limiter)

	MaxRequestsPerSecond = 5
)

func AllowRequest(host string) bool {
	mu.Lock()
	defer mu.Unlock()

	now := time.Now()
	l, exists := limiters[host]

	if !exists || now.After(l.windowEnd) {
		limiters[host] = &limiter{
			count:     1,
			windowEnd: now.Add(time.Second),
		}
		return true
	}

	if l.count >= MaxRequestsPerSecond {
		return false
	}

	l.count++
	return true
}
