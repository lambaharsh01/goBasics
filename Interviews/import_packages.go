package main

// WRONG IMPLEMENTATION
// is useless unless the package has an init() function that performs a side effect required by your program — usually registering itself somewhere

import (
	_ "fmt"
)

func BlankIdentifierImport() {
	println("Welcome")
	// println is a Built-in Go function

	// fmt.Println("Welcome") 
	// If you want to use Println, you must call it with the fmt package because 
	// 1. Part of the fmt package (requires import "fmt") 
	// 2. Fully formatted, reliable output
	// 3. Supports spacing, formatting, types, interfaces
}

// READ FOR CORRECT IMPLEMENTATION 

// Some packages run initialization code in their init() function.
// So sometimes you import them only to trigger their init logic, even if you don’t call anything from them.
// database driver does it to register itself inside the global database/db_name package

// EXAMPLES
// 1. Register a database driver: import _ "github.com/go-sql-driver/mysql"

// REGISTRATION GENERALLY LOOKS LIKE THIS:
// func init() {
//     sql.Register("mysql", &MySQLDriver{}) // registers itself globally
// }

// 2. Register an image format : import _ "image/png"

// WRONG IMPLEMENTATION
// is useless unless the package has an init() function that performs a side effect required by your program — usually registering itself somewhere
