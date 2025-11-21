package main

import (
	"fmt"
)


type Speaker interface {
	Speak()
}

type Person struct {
	Name string 
}

func(p Person) Speak() { // adding method to the struct
	fmt.Println("Person's Name is ", p.Name)
}

type Car struct {
	RegistrationNo int
}

func(c Car) Speak() {
	fmt.Println("Car's Registration no. is ", c.RegistrationNo)
}

func MakeSpeak(s Speaker) {
	s.Speak()
}

func InterfaceImplementation(){

	p := Person{Name:"Harsh"}
	c := Car{RegistrationNo:5664}

	MakeSpeak(p)
	MakeSpeak(c)

}

