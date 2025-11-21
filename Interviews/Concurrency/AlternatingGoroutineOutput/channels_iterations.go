package main

import (
	"fmt"
	"sync"
)

func WithChannelsInLoop() {

	n:=10

	c := make([]chan struct{}, n)

	for i := 0; i < n; i++ {
		c[i] = make(chan struct{})
	}

	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)


		go func(no int) {

			defer wg.Done()

			<-c[no]

			fmt.Println(no + 1)

			if no < n - 1 {
				c[no + 1] <- struct{}{}
			}

		}(i)
	}

	c[0] <- struct{}{}

	wg.Wait()
}
