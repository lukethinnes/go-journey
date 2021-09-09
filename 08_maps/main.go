package main

import "fmt"

func main() {
	//DEFINE MAP
	//emails := make(map[string]string)

	//DEFINE AND ADD KEY/VALUE PAIR
	emails := map[string]string{"Luke": "luke@luke.com", "Jeff": "jeff@jeff.com"}

	//emails["Luke"] = "lukethinnes@protonmail.com"
	//emails["Ceil"] = "xon@protonmail.com"

	//DELETE FROM MAP
	delete(emails, "Luke")
	fmt.Println(emails)

}
