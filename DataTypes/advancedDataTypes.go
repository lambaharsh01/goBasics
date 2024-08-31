package main

import "fmt"


//array //fixed arrays

//slices // dynamic arrays

// maps //objects learn more

//structs // structured objects

//pointers // new concept //Pointers are variables that store the memory address of another variable, allowing you to directly interact with memory in a more explicit way

// Interfaces //Interfaces in Go are used to define a set of method signatures that types (structs) must implement. They enable polymorphism in Go, allowing different types to be treated as the same type if they implement the same interface. // Interfaces do not create objects. They are purely an abstraction for behavior (i.e., methods).

// Channels //Channels are used for communication between goroutines

// What are goroutines?

func main(){

	// channelImplementation()
	// interfaceDeclaration()
	// pointerDeclarationWithPointer()
	// pointerDeclarationWithoutPointer()
	// structDeclaration()
	// mapsFunStaticDeclaration()
	// mapsFunDynamicDeclaration()
	// dynamicArray();
	// fixedArrays();
}



func channelImplementation(){

// INCORRECT
	// ch := make(chan int); //create a channel using the make function This creates a channel that can carry int values. Channels are of a specific type, and the type defines what kind of data can be sent through the channel.
	// ch <- 42; // send the value 42 into the channel
	// value :=ch;
	// fmt.Println(value);
	// close(ch);


// CORRECT LATER FIND OUT WHY IS THIS GIVING THIS BEHAVIOUR OF WHY ABOVE IS WRONG AND THIS IS RIGHT

ch := make(chan int) // Create a channel

go func() {
    ch <- 42 // Goroutine sending data
}()

value := <-ch; // Main goroutine receiving data
fmt.Println(value); // Output: 42


// Goroutines are a core feature of Go that facilitate concurrent programming. They are lightweight, managed by the Go runtime, and can be used to execute functions concurrently with minimal overhead. Goroutines, combined with channels, enable efficient and straightforward concurrent and parallel programming in Go.

}


// interfaces

func interfaceDeclaration(){

	d := Dog{}
    c := Cat{}
    
    MakeSound(d) // Outputs: Woof!
    MakeSound(c) // Outputs: Meow!

}

func MakeSound(a Animal) {
	fmt.Println(a.Speak())
}

// the interface is like a structure which a struct has to follow
// if the struct has the method Speak in this case 
// interface defines a contract that any type implementing it must have a Speak() method that returns a string. The key here is that any type which implements the Speak() method automatically satisfies the Animal interface
type Animal interface {
    Speak() string
}

type Dog struct{}

func (d Dog) Speak() string { // here d parameter should be an instance of the dog struct
    return "Woof!"
}

type Cat struct{}

func (c Cat) Speak() string {
    return "Meow!"
}


// interfaces end





func pointerDeclarationWithPointer(){
	var num int=2;
	p:=&num;
	modifyValueWithPointer(p);
	fmt.Println(num);
}
func modifyValueWithPointer(number *int) {
	*number = 100; // won't affect the original variable
	fmt.Println(*number);
}


func pointerDeclarationWithoutPointer(){
	var num int=2;
	modifyValueWithoutPointer(num);
	fmt.Println(num);
}
func modifyValueWithoutPointer(number int) {
	number = 100; // won't affect the original variable
	fmt.Println(number);
}




func structDeclaration(){
	
	type Person struct {
		Name string
		Age  int
	}
		
	var person1 Person = Person{Name: "Alice", Age: 30}
	fmt.Println(person1.Name);
}





func mapsFunStaticDeclaration(){

	studentGrades := map[string]int{
    "Alice": 91,
    "Bob": 81,
    "Charlie": 71,
	}

    fmt.Println(studentGrades);
    // fmt.Println(studentGrades["Alice"]);
}





func mapsFunDynamicDeclaration(){

	studentGrades := make(map[string]int)
	// Go maps are not safe for concurrent use without synchronization mechanisms like sync.Mutex or sync.Map.
	// it is like an object where values can be #hashed
    
	studentGrades["Alice"] = 90;
    studentGrades["Bob"] = 85;

    // fmt.Println(studentGrades);
    fmt.Println(studentGrades["Alice"]);
}





func dynamicArray(){
	var array []string = []string {"Hello", "World"}
	fmt.Println(array);
	
	array= append(array, "Go"); //push
	fmt.Println(array);
	
	array= array[:len(array)-2]; // pop
	fmt.Println(array);

}





func fixedArrays(){
	var array [3]int =[3]int {1,2,3};
	
	fmt.Println(array);
	fmt.Println(array[1]);
}