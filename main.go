package main

import "fmt"

func main() {

	fmt.Println("Ran")

	var slice []int

	fmt.Println(cap(slice), "--cap1")
	fmt.Println(len(slice), "--len1")
	
	for i:= 0; i<33; i++{
		slice = append(slice, i)
	} 

	fmt.Println(cap(slice), "--cap2")
	fmt.Println(len(slice), "--len2")

}