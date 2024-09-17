// methods are used to creat customised methods for structs/ other datatypes 
// in temrms of JS it is like creating an object class and then assigining a method inside of that class which applies the method's logic to the data attributes of the class instance Object ?

package main;
import "fmt";


type rectangle struct {
	height int
	width int
}

func (attrs rectangle) area() int {
 return attrs.height * attrs.width
}


func main(){
	var rectangle1 rectangle= rectangle{height:10, width:10};

	fmt.Println(rectangle1.area())
}