package main

import (
    "fmt"
)

func f(from string) { // we have a function which executes a loop from 0 to 499
    for i := 0; i < 100; i++ {
        fmt.Println(from, ":", i)
    }
}

func main() {

    fmt.Println("<--- FLAG 1 --->");

    go f("direct")
    go f("groutine")
    
    fmt.Println("<--- FLAG 2 --->");
    
    // fmt.Println("<--- FLAG 1 --->");

    // go f("direct")
    // f("direct")
    
    // fmt.Println("<--- FLAG 2 --->");

    // go f("goroutine")


    // time.Sleep(time.Second)
    fmt.Println("done");
    
}


// behavirour

// go is best for concrunt programming ( implementing multiple tasks at once on different worker threads to optimise available resources for resource intensive tasks especially when you have multi core CPU //dont be too reliant on that);

// now whenever we execute a normal code/ iteration/ recursion it is synchronously implemented on the main thread unless it is explisitly defined to be executed cuncurrently using go 


when i execute :{

    f("direct") //it first synchronously executes direct and once it is finesehed executing it moved on to executing goroutine

    go f("goroutine")

    // now if we just have this code without time.Sleep(time.Second) the program will exit without executing full worker thread because the program exits the execution if there is nothing on the main thread;

    // the result is somewhat similar to javascript one thing runs second thing runs and it does not have any task to run aftert that the program exits;

}


// when i execute :{
//         go f("direct") //when i do this bothe the functions run the iteration togeather on 2 sapeate threads 
//         go f("goroutine")
//         // the rsult is somewhat mix of both direct output and goroutine output
//     }



// when i execute :{

//         fmt.Println("<--- FLAG 1 --->");

//         go f("direct")
//         go f("direct")
        
//         fmt.Println("<--- FLAG 2 --->");
//         // the rsult is somewhat mix of both direct output and goroutine output
//     }


