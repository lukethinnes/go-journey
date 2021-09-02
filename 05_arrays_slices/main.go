package main

import "fmt"

func main() {
	//Arrays
	//var fruitArr [2]string

	//Assign values
	//fruitArr[0] = "Apple"
	//fruitArr[1] = "Orange"

	//Declare and assign
	//fruitArr := [2]string{"Apple", "Orange"}

	//fmt.Println(fruitArr)

	fruitSlice := []string{"Apple", "Orange", "Grapes"}

	fmt.Println(fruitSlice)

	//Wrapping in len() will produce number of values in array.
	fmt.Println(len(fruitSlice))
}
