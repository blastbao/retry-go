package retry

import (
	"time"
)

// ExponentialWaitStrategyRetry function. fn is the retry function.
func ExponentialWaitStrategyRetry(attempts int, sleep time.Duration, fn func() error) error {
	if err := fn(); err != nil {
		if s, ok := err.(stop); ok {
			return s.error
		}
		if attempts--; attempts > 0 {
			time.Sleep(sleep)
			return ExponentialWaitStrategyRetry(attempts, 2*sleep, fn)
		}
		return err
	}
	return nil
}

type stop struct {
	error
}