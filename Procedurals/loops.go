// There are 5 types of loops in goLang
// 1. forLoop
// 2. rangeLoop
// 3. whileLoop
// 4. infiniteLoop

package main;

import "fmt";

func main(){

	// forLoop(10);

	// fruitsSlice := []string{"apple", "banana", "cherry"}
	// rangeLoop(fruitsSlice);

	// whileLoop(6);


	// infiniteLoop();


}

func infiniteLoop(){
	for {
		fmt.Println("looping...")
		// use break to exit the loop
	}
}


func whileLoop(n int){

i := 0;

for i < n {
    fmt.Println(i)
    i++
}
}


func rangeLoop(fruits []string){

// for i, fruit := range fruits {
//     fmt.Println(i, fruit);
// }

// OR 

fruitsSlice := []string{"appla", "banana", "cherry"}

for i, fruit := range fruitsSlice {
    fmt.Println(i, fruit);
}

}


func forLoop(n int){

	for i := 0; i < n; i++ {

		if i==5 {
			continue;
		};
		
		fmt.Println(i, "from the for loop");

		if i==7 {
			break
		};
	}

}