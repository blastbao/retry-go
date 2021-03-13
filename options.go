package retry

import (
	"time"
)

const (
	DefaultMaxTries = 3
	DefaultTimeout  = 15 * time.Second
	DefaultDelayDuration    = 3 * time.Second
)

type retryOptions struct {
	maxTries        uint32
	timeout         time.Duration

	delayFunc   DelayFunc

	// 错误(码) 是否应该重试
	allowCodes []uint32				// P1
	denyCodes  []uint32				// P2
}

func newRetryOptions(options ...Option) *retryOptions {

	opts := &retryOptions{
		timeout:         DefaultTimeout,
		maxTries:        DefaultMaxTries,
		delayFunc: func(tries uint32) time.Duration { return DefaultDelayDuration },
	}

	for _, option := range options {
		option(opts)
	}

	return opts
}

type Option func(options *retryOptions)

// Timeout specifies the maximum time that should be used before aborting the retry loop.
// Note that this does not abort the operation in progress.
func WithTimeout(d time.Duration) Option {
	return func(options *retryOptions) {
		options.timeout = d
	}
}

// MaxTries specifies the maximum number of times op will be called by Do().
func WithMaxTries(tries uint32) Option {
	return func(options *retryOptions) {
		options.maxTries = tries
	}
}

func WithAllowCodes(codes ...uint32) Option {
	return func(options *retryOptions) {
		options.allowCodes = codes
	}
}

func WithDenyCodes(codes ...uint32) Option {
	return func(options *retryOptions) {
		options.denyCodes = codes
	}
}

func WithDelayFunc(df DelayFunc) Option {
	return func(options *retryOptions) {
		options.delayFunc = df
	}
}