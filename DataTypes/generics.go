// Generics is a feature that allows you to write code that can work with multiple types, without explicitly specifying the type. This means you can create functions or types that can be used with different data types, making your code more flexible and reusable.


package main

import "fmt"


func PrintValue[inputValue any](functionalVariable inputValue) {
	fmt.Println(functionalVariable)
}

func swapValues[inputValue any](firstFunctionalVariable, secondFunctionalVariable *inputValue) {
	// its very similar to [a,b]=[b,a] in javascript to swap the values of 2 variables
	*firstFunctionalVariable, *secondFunctionalVariable = *secondFunctionalVariable, *firstFunctionalVariable
}

// A generic struct to hold a value of any type
type Container[dynamicDataType any] struct {
	value dynamicDataType
}

func (containerStruct *Container[valueType]) GetValue() valueType {
	return containerStruct.value
}

func (containerStruct *Container[valueType]) SetValue(value valueType) {
	containerStruct.value = value
}

func main() {
	// Using the generic PrintValue function
	PrintValue("Hello, World!")  // string
	PrintValue(123)  // int
	PrintValue(true)  // bool

	// Using the generic Swap function
	a:= "Haa"
	b:= "Hee"
	swapValues(&a, &b)
	fmt.Println(a, b)

	// Using the generic Container struct
	container := Container[string]{value: "Hello world"} //Container[DATA_TYPE]{value: ACTUAL_VALUE} 
	fmt.Println(container.GetValue())  // 123
	container.SetValue("Bye World")
	fmt.Println(container.GetValue())  // 456
}