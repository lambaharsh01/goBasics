package main

import "fmt"

func Change(arr [2]string) {
	arr[1] = "World"
	// Arrays are passed by value, so only the local copy is modified.
}

// Arrays are copied when passed to a function, so changes don’t affect the original. Slices hold a pointer to their underlying array, so passing a slice automatically passes a reference, and changes affect the original without needing a pointer.
func ParametersArraysVsSlices() {

	arr := [2]string{}

	arr[0] = "Hello"
	Change(arr)

	fmt.Println(arr)
}
