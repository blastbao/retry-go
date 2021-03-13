package retry

import (
	"context"
	"math/rand"
	"time"
)

// BackoffFunc denotes a family of functions that control the backoff duration between call retries.
//
// They are called with an identifier of the attempt, and should return a time the system client should
// hold off for. If the time returned is longer than the `context.Context.Deadline` of the request
// the deadline of the request takes precedence and the wait will be interrupted before proceeding
// with the next iteration.
type DelayFunc func(attempt uint32) time.Duration

// BackoffFuncContext denotes a family of functions that control the backoff duration between call retries.
//
// They are called with an identifier of the attempt, and should return a time the system client should
// hold off for. If the time returned is longer than the `context.Context.Deadline` of the request
// the deadline of the request takes precedence and the wait will be interrupted before proceeding
// with the next iteration.
//
// The context can be used to extract request scoped metadata and context values.
type DelayFuncContext func(ctx context.Context, attempt uint) time.Duration




// DelayConstant is very simple: it waits for a fixed period of time between calls.
func DelayConst(delay time.Duration) DelayFunc {
	return func(attempt uint32) time.Duration {
		return delay
	}
}

// DelayJitter adds random jitter to constant delay.
//
// For example wait=1s and jitter=0.10 can generate waits between 900ms and 1100ms.
func DelayConstJitter(delay time.Duration, jitter float64) DelayFunc {
	return func(attempt uint32) time.Duration {
		return JitterUp(delay, jitter)
	}
}

// DelayExpBackoff produces increasing intervals for each attempt.
//
// The scalar is multiplied times 2 raised to the current attempt. So the first
// retry with a scalar of 100ms is 100ms, while the 5th attempt would be 1.6s.
func DelayExpBackoff(scalar time.Duration) DelayFunc {
	return func(attempt uint32) time.Duration {
		return scalar * time.Duration(ExponentBase2(attempt))
	}
}

// DelayExpBackoffJitter creates an exponential backoff like DelayExpBackoff does, but adds jitter.
func DelayExpBackoffJitter(scalar time.Duration, jitter float64) DelayFunc {
	return func(attempt uint32) time.Duration {
		return JitterUp(scalar * time.Duration(ExponentBase2(attempt)), jitter)
	}
}

// JitterUp adds random jitter to the duration.
//
// This adds or subtracts time from the duration within a given jitter fraction.
// For example for 10s and jitter 0.1, it will return a time within [9s, 11s])
func JitterUp(duration time.Duration, jitter float64) time.Duration {
	multiplier := jitter * (rand.Float64()*2 - 1)
	return time.Duration(float64(duration) * (1 + multiplier))
}

// ExponentBase2 computes 2^(a-1) where a >= 1. If a is 0, the result is 0.
func ExponentBase2(a uint32) uint {
	return (1 << a) >> 1
}