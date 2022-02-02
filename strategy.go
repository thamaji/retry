package retry

import (
	"math/rand"
	"time"
)

type Strategy interface {
	Delay(int) time.Duration
	Max() int
}

var DefaultStrategy = ExponentialBackoffAndEqualJitter{
	MaxRetryCount: 10,
	BaseDelayTime: 1 * time.Second,
	MaxDelayTime:  5 * time.Minute,
}

func MaxRetryCount(c int) Strategy {
	s := DefaultStrategy
	s.MaxRetryCount = c
	return s
}

type ExponentialBackoff struct {
	MaxRetryCount int
	BaseDelayTime time.Duration
	MaxDelayTime  time.Duration
}

func (strategy ExponentialBackoff) Max() int {
	return strategy.MaxRetryCount
}

func (strategy ExponentialBackoff) Delay(c int) time.Duration {
	delay := strategy.BaseDelayTime * 2 << c
	if delay > strategy.MaxDelayTime {
		delay = strategy.MaxDelayTime
	}
	return delay
}

type ExponentialBackoffAndFullJitter struct {
	MaxRetryCount int
	BaseDelayTime time.Duration
	MaxDelayTime  time.Duration
}

func (strategy ExponentialBackoffAndFullJitter) Max() int {
	return strategy.MaxRetryCount
}

func (strategy ExponentialBackoffAndFullJitter) Delay(c int) time.Duration {
	delay := strategy.BaseDelayTime * 2 << c
	if delay > strategy.MaxDelayTime {
		delay = strategy.MaxDelayTime
	}
	return time.Duration(rand.Int63n(int64(delay)))
}

type ExponentialBackoffAndEqualJitter struct {
	MaxRetryCount int
	BaseDelayTime time.Duration
	MaxDelayTime  time.Duration
}

func (strategy ExponentialBackoffAndEqualJitter) Max() int {
	return strategy.MaxRetryCount
}

func (strategy ExponentialBackoffAndEqualJitter) Delay(c int) time.Duration {
	delay := strategy.BaseDelayTime * 2 << c
	if delay > strategy.MaxDelayTime {
		delay = strategy.MaxDelayTime
	}
	return delay/2 + time.Duration(rand.Int63n(int64(delay/2)))
}
