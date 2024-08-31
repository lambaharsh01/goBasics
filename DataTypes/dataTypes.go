package main

import "fmt"

// there are 7 datatypes in go.
// bool 
// int
// float32, float64
// complex64, complex124
// string
// rune -- a rune is a data type that represents a Unicode code point. It is an alias for the int32 type and is used to handle individual characters in Unicode. 
// array var numbers [5]int if the value of the elements is fixed 
// else use slice var numbers []int without a value in [];
// struct
// map
// Channel chan


// Special Data Types
// Interface




	// var isActive bool= true; || isActive:= false; -assignment;
	// isActive=true; -reassignment;

	// var greeting rune = 'H'
	// fmt.Println(greeting); //72 in int
	


func main(){

	// var numbers [5]int = [5]int{1, 2, 3, 4, 5};

	// fmt.Println(numbers);

	// var colors []string = []string{"red", "green", "blue"};
	// fmt.Println(len(colors)); // len(length) and cap(capacity) tell the length and the number of elements inside of an array/slice
	// fmt.Println(cap(colors));


	// Slice could be appended pushed 
	// colors=append(colors, "orange"); // push
	// fmt.Println(colors);
	
	// colors = colors[:len(colors)-1]; //create a new slice from the original slice colors, excluding its last element. 
	// colors = colors[:len(colors)-2]; // excluding last 2 elements
	// fmt.Println(colors);

	// fmt.Println(colors[2]);




	// OBJECTS

	
type Person struct {
    Name string
    Age  int
}
	
	var p Person = Person{Name: "Alice", Age: 30}
	fmt.Println(p.Name, p.Age);

}