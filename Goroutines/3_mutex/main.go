package main

import (
	"fmt"
	"sync"
)

// 1. sync.Mutex is used to protect shared resources from concurrent access.
// 2. Use .Lock() to enter the critical section and .Unlock() to leave it.
// 3. Forgetting to unlock will cause deadlocks.

func main(){
	
	var commonInt int
	var mx sync.Mutex 
	
	var waitGroup sync.WaitGroup

	for i:=0; i<10000; i++ {
		
		waitGroup.Add(1)

		go func(waitGroup *sync.WaitGroup, i int){
			defer waitGroup.Done()

			mx.Lock()
			
			commonInt++

			mx.Unlock()

		}(&waitGroup, i)
	}
	
	waitGroup.Wait()

	fmt.Println("Total Count is ", commonInt)
	
}









// 🔐 sync.RWMutex:


	// var mu sync.RWMutex
	// var count int

	// // Reader
	// go func() {
	// 	mu.RLock()
	// 	fmt.Println("Reading count:", count)
	// 	mu.RUnlock()
	// }()

	// // Writer
	// go func() {
	// 	mu.Lock()
	// 	count++
	// 	mu.Unlock()
	// }()
	

// It is a type of mutex, but more feature-rich than a basic sync.Mutex.
// It does not restrict a goroutine to write only once — instead:
// Readers (RLock) can access the data simultaneously if no writer is holding the lock.
// Writer (Lock) gets exclusive access, i.e., no other goroutine (reader or writer) can access the resource while the writer is active.
// Once a writer is done, it can Unlock, and any goroutine (reader or writer) can acquire the lock again.
// Deadlocks can still happen if you don’t release the locks properly (e.g., forget to call Unlock()).
// It's for synchronizing access to shared data when there are both readers and writers, not one-time-only access.





// ⚙️ sync/atomic:


	// import "sync/atomic"

	// var count int64

	// go func() {
	//     atomic.AddInt64(&count, 1)
	// }()

	// go func() {
	//     fmt.Println("Count:", atomic.LoadInt64(&count))
	// }()


// Not a mutex at all — it’s a lock-free alternative for very specific cases (simple shared variables like integers or booleans).
// It uses atomic CPU instructions to make sure no two goroutines corrupt a variable when reading/updating it.
// It's very fast and efficient but only works on basic types.
// Since it's lock-free, there are no deadlocks (you’re right about that!), but you still need to design your logic carefully to avoid race conditions or logic bugs.

