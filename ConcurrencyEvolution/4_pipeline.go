// PIPELINE is about staged transformation, not parallelism.
// It does not require multiple instances of a function running concurrently.
// In Go, pipelines are typically structured so that each stage processes
// items as soon as they are produced by the previous stage.
// Data flows incrementally between stages, allowing downstream stages
// to begin work without waiting for the entire upstream stage to complete.
// This enables streaming and overlap between stages.

// IMPROVENEMTNATS in the following code (Depends on the situation)
// 1. Blocking due to UNBUFFERED CHANNELS: this makes the pipeline more synchronized and reduces parallel overlap between stages. adding BUFFERED CHANNELS allows greater concurrency and helps HANDLES SMALL BURSTS of data.

// 2. the returned channels can be named out := make(chanel int) // because function name describes what they do

package main

import "fmt"

func Pipeline() {

	stage1 := GenerateNumber(10)
	stage2 := FilterEven(stage1)
	stage3 := Square(stage2)
	stage4 := StrConv(stage3)

	for res := range stage4 {
		fmt.Println(res)
	}

}

func GenerateNumber(n int) <-chan int {
	nums := make(chan int)

	go func() {
		defer close(nums)
		for i := 1; i <= n; i++ {
			nums <- i
		}
	}()

	return nums
}

func FilterEven(nums <-chan int) <-chan int {
	even := make(chan int)

	go func() {
		defer close(even)
		for num := range nums {
			if num%2 == 0 {
				even <- num
			}
		}
	}()

	return even
}

func Square(nums <-chan int) <-chan int {
	squares := make(chan int)

	go func() {
		defer close(squares)
		for num := range nums {
			squares <- num * num
		}
	}()

	return squares
}

func StrConv(nums <-chan int) <-chan string {
	strs := make(chan string)

	go func() {
		defer close(strs)
		for num := range nums {
			strs <- fmt.Sprintf("Result:%d", num)
		}
	}()

	return strs
}
