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
