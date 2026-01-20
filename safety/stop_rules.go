package safety

type Signal string

const (
	SignalWAF        Signal = "WAF_DETECTED"
	SignalAuthLock   Signal = "AUTH_LOCKOUT"
	SignalErrorSpike Signal = "ERROR_SPIKE"
)

func ShouldStop(signals []Signal) bool {
	for _, s := range signals {
		switch s {
		case SignalWAF, SignalAuthLock:
			return true
		}
	}
	return false
}
