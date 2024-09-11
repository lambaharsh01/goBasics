// In the Go , when a function might run into a problem or fail to do its job, it doesn't hide this fact or use special mechanisms to signal the error. Instead, it simply returns an extra value along with its regular output. This extra value is specifically for indicating whether an error occurred.

package main;

import ("fmt"; "errors");



func divide(a, b float64) (float64, error) {

	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	return a / b, nil
}

func main() {
	result, err := divide(10.4, 0)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}


// What is Error and Panic in go ?

// if("error"){

// 1. Errors are part of the normal execution flow and are returned as values.

// 2. They occur in expected situations, such as a failed API call, a file not found, or invalid input.

// 3. Go functions often return both a result and an error. You check if the error is nil to determine if everything went fine.

// }else{

// 1. Panic is a mechanism used when something goes catastrophically wrong, and the program cannot continue running in a normal way.

// 2. When you call panic(), the program immediately stops executing the normal flow and starts unwinding the stack, running any deferred functions.

// 3. When you call panic(), the program immediately stops executing the normal flow and starts unwinding the stack, running any deferred functions.


// }