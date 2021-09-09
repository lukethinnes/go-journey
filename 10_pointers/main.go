package main

import "fmt"

func main() {
	//Pointers point to the memorry address of a value in a variable.

	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", b)

	//Use * to read the addresses value
	fmt.Println(*b)
	fmt.Println(*&a)

	//Change value with a pointer
	*b = 10
	fmt.Println(a)
}
