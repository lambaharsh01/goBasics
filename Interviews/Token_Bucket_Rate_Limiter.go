package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

type TokenDetails struct {
	Cap        int
	Tokens     int
	Rate       int
	LastRefill time.Time
	Mx         sync.Mutex
}

func (td *TokenDetails) Allow() bool {
	td.Mx.Lock()
	defer td.Mx.Unlock()

	now := time.Now()
	td.Tokens += int(now.Sub(td.LastRefill).Seconds()) * td.Rate
	td.Tokens = int(math.Min(float64(td.Tokens), float64(td.Cap)))
	td.LastRefill = now

	if td.Tokens > 0 {
		td.Tokens--
		return true
	}
	return false
}
func Constructor(cap, rate int) *TokenDetails {
	return &TokenDetails{
		Cap:        cap,
		Tokens:     cap,
		Rate:       rate, // token per second
		LastRefill: time.Now(),
	}
}
func main() {

	m := map[int]*TokenDetails{}
	var mx sync.Mutex

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			mx.Lock()
			t, ok := m[1]
			if !ok {
				t = Constructor(5, 3)
				m[1] = t
			}
			mx.Unlock()

			if t.Allow() {
				fmt.Println("Allowed")
			} else {
				fmt.Println("Not Allowed")
				time.Sleep(time.Millisecond * 1100)
			}
		}()

	}

	wg.Wait()
}
