package main

import (
	"fmt"
	"sync"
)

func FanIn() {
	// Multiple Producers Single Consumer

	ch := make(chan int)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			ch <- num
		}(i)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for curr := range ch {
		fmt.Println(curr)
	}
}
