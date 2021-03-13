package retry

import (
	"context"
	"testing"
	"time"
)

func TestRetry_Base(t *testing.T) {

	err := DoCtx(
		context.Background(),
		func(_ context.Context) error {
			return nil
		},
		WithTimeout(10*time.Second),
		WithMaxTries(6),
		WithDelayFunc(DelayConst(1)),
	)

	_ = err
}