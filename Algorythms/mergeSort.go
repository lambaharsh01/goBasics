package main

import "fmt"

func main(){ 

	var unsortedArray []int = []int {84, 1, 55, 24, 4, 12, 94, 33, 16};
	fmt.Println(devide(unsortedArray));

}

func devide(array []int) []int {

	if(len(array)<2){
		return array;
	}

	var middleIndex int= len(array)/2;  //Go performs integer division when both operands are integers. In Go, dividing two integers results in an integer, and the fractional part is truncated (not rounded).

	// slicing :number will lead to the slice of the right side of the number
	// slicing number: will lead to the slice of the left side of the number

	var firstHalf []int= array[:middleIndex];
	var secondHalf []int= array[middleIndex:];

	return merge(devide(firstHalf), devide(secondHalf) )

}


func merge(array1 []int, array2 []int) []int {

	var leftIndex, rightIndex int;

	var sortedArray []int;

	for leftIndex < len(array1) && rightIndex < len(array2){

		if array1[leftIndex] < array2[rightIndex]{
			sortedArray= append(sortedArray, array1[leftIndex]);
			leftIndex++;
		}else{
			sortedArray= append(sortedArray, array2[rightIndex]);
			rightIndex++;
		}
	}

	sortedArray=append(sortedArray, array1[leftIndex:]...);

	return append(sortedArray, array2[rightIndex:]...);
}





