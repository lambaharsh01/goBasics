package main

import (
	"fmt"
	"sync"
)


func main() {
	var intPrint bool
	var mx sync.Mutex
	var wg sync.WaitGroup
	wg.Add(2)

	var iterations int

	go func(){
		defer wg.Done()
		for i:= 'A'; i<='E'; {
			iterations++
			mx.Lock()
				if !intPrint {	
					fmt.Print(string(i))
					intPrint = true
					i++
				}
			mx.Unlock()
		}
	}()

	go func(){
		defer wg.Done()
		for i:=1; i <=5; {
			iterations++
			mx.Lock()
				if intPrint {
					fmt.Print(i)
					intPrint = false
					i++
				} 
			mx.Unlock()
		}
	}()

	wg.Wait()

	fmt.Println()
	fmt.Println(iterations, "--iterations")


	// more than 5000 iteration for a scenario where n = 10 
	// this condition is called BUSY WAITING 
	// Busy waiting (also called spin waiting or a spinlock) is when a goroutine or thread keeps actively checking for a condition to become true — instead of sleeping or blocking — and in the process, uses CPU cycles doing nothing productive.
}

// func main() {

// 	chanOdd := make(chan bool)
// 	chanEven := make(chan bool)

// 	var wg sync.WaitGroup

// 	wg.Add(2)

// 	go func(){
// 		wg.Done()
// 		for i:= 1; i<= 10; i+=2 {
// 			<-chanOdd
// 			fmt.Print(" ",i)
// 			chanEven <- true
// 		}
// 	}()

// 	go func() {
// 		defer wg.Done()

// 		chanOdd <- true

// 		for i:=2; i<=10; i+=2 {
// 			<-chanEven
// 			fmt.Print(" ",i)
// 			if i < 10 {
// 				chanOdd <- true
// 			}
// 		}
// 	}()

// 	wg.Wait()

// }

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {

// 	chanChar := make(chan bool)
// 	chanInt := make(chan bool)

// 	var wg sync.WaitGroup

// 	wg.Add(2)

// 	go func(){
// 		wg.Done()
// 		for i:= 'A'; i<='E'; i++ {
// 			<-chanChar
// 			fmt.Print(string(i))
// 			chanInt <- true
// 		}
// 	}()

// 	go func() {
// 		defer wg.Done()
// 		for i:=1; i<=5; i++ {
// 			<-chanInt
// 			fmt.Print(i)
// 			if i < 5{
// 				chanChar <- true
// 			}
// 		}
// 	}()

// 	chanChar <- true

// 	wg.Wait()

// }

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main(){

// 	var wg sync.WaitGroup

// 	charChan := make(chan bool)
// 	intChan := make(chan bool)

// 	wg.Add(2)

// 	go func(){

// 		defer wg.Done()
// 		for i:= 'A'; i <= 'E'; i++ {
// 			<-charChan
// 			fmt.Print(string(i))
// 			intChan <- true
// 		}

// 		}()

// 	go func() {
// 		defer wg.Done()

// 		for i:=1; i<=5; i++ {

// 			<-intChan
// 			fmt.Print(i)

// 			if i < 5 {
// 				charChan <- true
// 			}
// 		}
// 		}()

// 		charChan <- true
// 	wg.Wait()
// }

// // package main

// // import "fmt"

// // // func main() {

// // // 	text := " this      is   string. "
// // // 	// " this is string. "

// // // 	var chars string

// // // 	for i, r := range text {

// // // 		if i != 0 && text[i] == ' ' && text[i-1] == ' ' {
// // // 			continue
// // // 		}
// // // 		chars += string(r)
// // // 	}

// // // 	fmt.Println(chars)

// // // 	// for i, c := range chars {
// // // 	// 	if i!= 0 && c != ''

// // // 	// }

// // // 	// fmt.Println(strings.ReplaceAll(text, "  ", " "))
// // // }

// // func main() {

// // 	arr := [5]int{1, 2, 3, 4, 5}
// // 	slice := arr[:3]
// // 	// 0, 1

// // 	slice = append(slice, 500, 600)
// // 	// 0, 1, 500, 600

// // 	slice[2] = 1000
// // 	// 0, 1, 1000, 600

// // 	fmt.Printf("arr: %v\n", arr)
// // 	fmt.Printf("slice: %v\n", slice)
// // }

// // // ACTUAL OUTPUT
// // // arr: [1 2 1000 500 600]
// // // slice: [1 2 1000 500 600]
