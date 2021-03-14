package retry

import (
	"time"
)

const (
	DefaultMaxTries = 3
	DefaultDelayDuration    = 3 * time.Second
)

type retryOptions struct {
	maxTries        uint32
	delayFunc   DelayFunc

	// 错误(码) 是否应该重试
	allowCodes []uint32				// P1
	denyCodes  []uint32				// P2
}

func newRetryOptions(options ...Option) *retryOptions {

	opts := &retryOptions{
		maxTries:        DefaultMaxTries,
		delayFunc: func(tries uint32) time.Duration { return DefaultDelayDuration },
	}

	for _, option := range options {
		option(opts)
	}

	return opts
}

type Option func(options *retryOptions)

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