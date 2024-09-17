// timers and tickers are used for handling time-related operations. 
// timer is like setTimeout functio in javascript
// timers are a blocking scope in go when in javascript setTimout is a non blocking scope
package main;

import (
	"fmt"
	"time"
)

func main(){
	// timers()
	ticker()
}


func timers(){

	timer := time.NewTimer(2 * time.Second)

	fmt.Println("Start")

	<-timer.C  // Block until the timer fires
	fmt.Println("Timer fired!")
}


func ticker(){
// A ticker is similar to a timer, but it repeatedly sends the current time on its channel at regular intervals.

ticker := time.NewTicker(1 * time.Second)

go func() { // putting it outside the groutine will make it tick until the program is closed manually
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	// Stop the ticker after 5 seconds
	time.Sleep(5* time.Second)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}