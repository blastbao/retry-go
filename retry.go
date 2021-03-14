package retry

import (
	"context"
	"time"
)


func Do(fn func() error, retryOptions ...Option) error {
	return DoCtx(context.Background(), func(_ context.Context) error {
		return fn()
	}, retryOptions...)
}

func DoCtx(ctx context.Context, fn func(context.Context) error, retryOptions ...Option) error {
	opts := newRetryOptions(retryOptions...)
	attempts := uint32(0)
	for {
		err := fn(ctx)
		if err == nil { // no error, return directly
			return nil
		} else if e, ok := err.(Err); ok { // if err is Err type, then check code whitelist
			if len(opts.allowCodes) != 0 {
				if !Contains(opts.allowCodes, e.Code) {
					return err
				}
			} else if len(opts.denyCodes) != 0 {
				if Contains(opts.denyCodes, e.Code) {
					return err
				}
			}
		} else { // if err is not type Err, just retry
		}

		attempts++
		if attempts >= opts.maxTries { // reach limit
			return err
		}

		// wait for the next duration or until the context is done, whichever comes first
		t := time.NewTimer(opts.delayFunc(attempts))
		select {
		case <-t.C: // delay duration elapsed, continue loop and retry
		case <-ctx.Done():
			// context cancelled, kill the timer if it hasn't fired, and return the last error we got
			if !t.Stop() {
				<-t.C
			}
			return err
		}
	}

}