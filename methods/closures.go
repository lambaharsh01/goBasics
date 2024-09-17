package main;

import "fmt";

func closureFunc() func(int) int{
	var sum int= 0;

	return func(x int) int{
		sum+= x;
		return sum;
	}
}

func main(){

	initial := closureFunc();

	fmt.Println(initial(1));
	fmt.Println(initial(1));
	fmt.Println(initial(1));
	fmt.Println(initial(1));

}