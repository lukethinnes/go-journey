package main

import "fmt"

func main() {
	//var name = "Luke"

	name := "Luke"
	var age int32 = 27
	const isAlive = true

	fmt.Println(name, age, isAlive)
	fmt.Printf("%T\n", isAlive)
}
