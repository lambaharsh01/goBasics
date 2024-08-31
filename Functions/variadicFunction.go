// variadic functions are functions that can take a variable number of arguments. This allows you to pass zero or more arguments of the same type to the function. Variadic functions use the ... syntax to indicate that the function accepts a varying number of arguments

package main;

import "fmt";

func main(){

	fmt.Println(sum(1,2,3,4,5)) ; //accepts n number of arguements

}


func sum(nums ...int) int { // this ) int { because this function return a number;
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}