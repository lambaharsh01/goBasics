package main;

import "fmt";

func divide(a, b int) (int, error) {
	// a multiple return function is a function that can return more than one value at the same time. This feature is useful when you want to return a result along with additional information, such as an error message, without having to rely on exceptions or side effects.
    if b == 0 { 
		return 0, fmt.Errorf("division by zero is not allowed") 
	}

    return a / b, nil ;
}

func main() {
    result, err := divide(0, 1);
    
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result)
    }

	// OR 

	var value1, value2 int =getValues();
	fmt.Println(value1, value2);
}


func getValues()(int, int){
	return 1,2
}