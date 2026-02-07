package main

import (
	"sync"
	"time"
)

type TokenBucket struct {
	Cap        float64
	Tokens     float64
	RPS        float64
	LastRefill time.Time
	Mx         sync.Mutex
}

func (t *TokenBucket) Allow() bool {
	t.Mx.Lock()
	defer t.Mx.Unlock()

	now := time.Now()
	secDiff := float64(now.Sub(t.LastRefill).Seconds()) * t.RPS
	t.Tokens += secDiff   // ellipse has to be added to the existing tokens
	if t.Cap < t.Tokens { // if tokens(because of the time passes) are more than the capacity we cap it
		t.Tokens = t.Cap
	}

	t.LastRefill = now

	if t.Tokens >= 1 { // if 0.2 and 0.2-1 it becomes -0.8 breaks the algo
		t.Tokens--
		return true
	}

	return false
}

func Constructor(cap, rps float64) *TokenBucket {
	return &TokenBucket{
		Cap:        cap,
		Tokens:     cap,
		RPS:        rps,
		LastRefill: time.Now(),
	}
}

func main() {
// Implement it here
}
