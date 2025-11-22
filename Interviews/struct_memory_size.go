package main

import "unsafe"

type Stud struct {
	x int64 
	y int64
	// z int
	// w int
}

func GoStructMemorySize(){

	// On 64-bit systems (most common): 8 bytes
	// On 32-bit systems: 4 bytes

	a := Stud{x:1}
	// b := Stud{x:1, y: 10000,z:10, }
	// c := Stud{x:93239, y:45665, z:2323}

	println(unsafe.Sizeof(a)) 
	// println(unsafe.Sizeof(b)) 
	// println(unsafe.Sizeof(c)) 

}
