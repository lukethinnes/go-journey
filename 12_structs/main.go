package main

import (
	"fmt"
	"strconv"
)

//Person defined with struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

//Methods

//Greeting method (VALUE RECEIVER)

//Requires identifier then name of struct

func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

//hasBirthday method will be a pointer receiver

func (p *Person) hasBirthday() {
	p.age++
}

func main() {
	//Initialize a person using struct
	person1 := Person{
		firstName: "Luke",
		lastName:  "Thinnes",
		city:      "Denver",
		gender:    "M",
		age:       27,
	}
	fmt.Println(person1)

	//Get a single field using dot notation...
	person1.hasBirthday()
	fmt.Println(person1.greet())
}
