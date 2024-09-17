// In Go, a panic is a built-in mechanism that represents an unexpected error or condition during the execution of a program. When a program panics, it stops executing the normal flow and begins to unwind the stack of function calls, which eventually leads to the termination of the program unless the panic is recovered.

package main;

import ("fmt");


func main(){
	// panic();
	recoverFromPanic()
}

func panicFunction(){
	
		fmt.Println("hello world")
	
		panic("a problem")

}

func recoverFromPanic(){
	defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from:", r)
        }
    }();

	fmt.Println("First Execution");

	panic("Some Problem");
}