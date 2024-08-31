// Go does not have ternary operators instead it enchorates if-else;

package main;

import "fmt";

func main(){
	switchCase(5)
	// ifElse(5)
}

func switchCase(number int){

	// if the condition has to be boolean then there is no need to take an switch arguement

	// switch number{
	// case number > 5 :
	// 	fmt.Println("number is greater than 5")		
	// case number == 5 :
	// 	fmt.Println("number is equal to 5")
	// default:
	// 	fmt.Println("number is less than 5")
	// }


	// if the case has to be as per the input then there should be the need for swith variable 
	switch number{
	case 5 :
		fmt.Println("number is 5")
	default:
		fmt.Println("number is not 5 5")
	}

}


func ifElse(number int){
// the statements should always be in bool (not truly or falsy values like in javascript)
    if number > 5 {
        fmt.Println("number is greater than 5")
    } else if number == 5 {
        fmt.Println("number is equal to 5")
    } else {
        fmt.Println("number is less than 5")
    }
}