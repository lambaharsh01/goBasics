// a function can not be declaed inside of another function in go;
// functions should be independent units;

package main;
import "fmt";

func main(){
	someFunction();
}


func someFunction(){
	fmt.Println("Function Invoked");
}