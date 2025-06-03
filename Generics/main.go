package main

import "fmt"

func main() {

	strArr := []string{"A", "B", "A", "A", "A", "A", "A", "A", "A"}
	fmt.Println(IndexOf(strArr, "B"))

	intArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 0, 0, 0, 1}
	fmt.Println(IndexOf(intArr, 6))

	fmt.Println(Add(1,2.3)) // float and int are not that different and first arg is implicitly converted into float in this example

}

// WHEN IT COMES TO GENERICS they're there so you can have a single function for similar operations on different data types
// To make go less verbose 

// T stands for the custom data types which are acceptable which would have been 1 datatype in normal scenario
// Now Considering T could be of multiple datatype but if T is passed as the datatype to multiple arguments it has to remail consistent across all the T used in the generic call 

// if my Generic Says IndexOf[T int | float64 | string] (arr []T, val T)
// if i init with a string arr the val arg should also be of the same time as first arg 
// cant have IndexOf([]string, int) has to same data type IndexOf([]string, string) or IndexOf([]int, int) ...and so on.


func IndexOf[T int | float64 | string] (arr []T, val T) int {
	for i, elem := range arr{
		if elem == val {
			return i
		}
	}
	return -1
}

func Add[T int | float64] (a, b T) T {
	return a+b
}