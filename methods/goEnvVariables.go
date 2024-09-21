
// first initialsie a go.mod file by using the command 
//go mod init "pathNameYouWantFOrYourModules" 

//It performs two main actions:
// 1. Creates the go.mod File: The go mod init command initializes a new Go module by creating a go.mod file in the current directory. This file is used to track the dependencies of your Go project and manage its module settings.
//like {// module envExample
// go 1.20
//}

// 2.Defines the Module Path: like require github.com/joho/godotenv v1.5.1 // indirect


// ** IMPORTANT **
// The envExample name serves as the module path or namespace for your project and its packages. When you build larger projects and create packages within this module, the envExample name will be part of the import path.

// If you add a package config inside the envExample module, you would import it as:
// import "envExample/config"
// When sharing this module or using it in other projects, you would refer to it by the module name envExample (or the full repository URL if published).

// The go.mod file tracks the external libraries (modules) your project depends on. This makes it easier to manage, upgrade, and keep track of third-party packages. For example, when you run go get, the required modules are added to your go.mod file.
// The go.mod file is the root of your Go module and will manage your project's dependencies



// STEPS 
// 1. cmd // go mod init "pathNameYouWantFOrYourModules" 
// 2. create a go.sum file in the directory
// 3. install packages //go get github.com/joho/godotenv@latest
// all set and youre good to use the package 

package main

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
)

func main() {


	goDotEnv();

	return;


	// Get the value of the environment variable
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")

	// Check if the variable is not set
	if dbUser == "" || dbPass == "" {
		fmt.Println("Environment variables DB_USER or DB_PASS are not set.")
		return
	}

	fmt.Printf("Database User: %s\n", dbUser)
	fmt.Printf("Database Password: %s\n", dbPass)
}


func goDotEnv(){
	err := godotenv.Load()
    if err != nil {
        fmt.Println("Error loading .env file")
    }

    dbUser := os.Getenv("DB_USER")
    dbPass := os.Getenv("DB_PASS")

    fmt.Printf("Database User: %s\n", dbUser)
    fmt.Printf("Database Password: %s\n", dbPass)
}