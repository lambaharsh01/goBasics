package main

import (
	"fmt"
)

func VariableDeclaration() {
	
	var x, y, z = 1, 2.4, "welcome"
	// x, y, z := 1, 2.4, "welcome"

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	strArr := [2]string{} // initializes and empty ["", ""] string array implicitly
	strArr[0]="Harsh"
	strArr[1]="Lamba"

	fmt.Println(strArr)

}


func ZeroValueOfEmptyInterface(){

	var anyVar interface{} // OR any
	
	if anyVar == nil {
		println("well it's nil alright")
	} else {
		println("naa not nil")
	}

	// 1. Can only perform assignment = 
	// 2. Comparison == operations can be done 
	// 3. No arithmetic operations 
	// 4. Can't access fields or methods a.Name  or a[0]
	// 5. Can't pass it to a function expecting a specific type // fmt.Println(len(a))

	// 6. TYPE SWITCH CAN BE DONE 
	switch v := a.(type) {
	case int:
		fmt.Println("int:", v)
	case string:
		fmt.Println("string:", v)
	default:
		fmt.Println("unknown type")
	}


}


func ArrayLiteralEllipsisLength(){
	arr := [...]int{1, 2, 3, 4} // array is defined as  [4]int
	println(arr) 
}
