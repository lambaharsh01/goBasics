package ConcurrencyPatterns

import (
	"fmt"
	"sync"
	"time"
)

type Token struct {
	Cap float64
	RefreshRate float64
	Token float64
	LastRequestAt time.Time
	Mu sync.Mutex
}

func GenerateToken(cap, refreshRate float64) *Token {
	return &Token{
		Cap: cap,
		RefreshRate: refreshRate,
		Token: cap,
		LastRequestAt: time.Now(),
	}
}

func (t *Token) Allow() bool {

	t.Mu.Lock()
	defer t.Mu.Unlock()

	currTime := time.Now()
	// how much time has passed since the last refill (in seconds) aka elapsed
	timePassedSinceLastRefill := currTime.Sub(t.LastRequestAt).Seconds()
	t.Token += timePassedSinceLastRefill * t.RefreshRate

	// add tokens based on time passed: newTokens = elapsedSeconds * refreshRate
	// now the t.Token might be more than 5 because if the user did not request for 40 seconds so 40 * 1 (refresh rate of 1 token generated per second) = 40 token so with this if clause we bring it down to it's capacity which is 5
	if t.Token > t.Cap {
		t.Token = t.Cap
	}

	t.LastRequestAt = currTime

	// now the no. of token reduced by 1 if there are tokens available representing that 1 request has been processed and update the tokens available
	if t.Token >= 1 {
		t.Token = t.Token - 1
		return true
	}

	// if no tokens are available we limit the request
	return false
}


func SingleTokenBucket() {
	token := GenerateToken(5, 1)

	for i:=0; i < 20; i++ {
		if token.Allow() {
			fmt.Println("Allowed", i)
		} else {
			fmt.Println("Blocked", i)
		}

		time.Sleep(300 * time.Millisecond)
	}
}

func MultiTokenBucket() {

	var mx sync.Mutex
	var users = make(map[int]*Token)
	
	for i:= 0; i < 20; i++ {
		mx.Lock()

		userID := i%2

		// Create token bucket for the user if it doesn't exist
		if _, ok := users[userID]; !ok {
			users[userID] = GenerateToken(5,2)
		}
		
		// Copy pointer so we can safely use it after unlocking the map
		userToken := users[userID]
		mx.Unlock()

		// Unlock here because each Token has its own internal mutex.
		// The map-level mutex should ONLY protect map reads/writes.
		// After retrieving the pointer, the Token’s own mutex ensures thread-safe rate limiting without blocking the entire users map.

		if userToken.Allow() {
			fmt.Println("User:", userID, "allowed")
		} else {
			fmt.Println("User:", userID, "Blocked")
		}

		time.Sleep(100 * time.Millisecond)
	} 
	
}


func RateLimiter() {

	// “Lets tokens grow over time and spends them per request. Allows bursts.”
	// “Queue where requests leak at a fixed rate. Smooth output but no bursts.”

	// We are using Token Bucket

	// SingleTokenBucket()
	MultiTokenBucket()
}
