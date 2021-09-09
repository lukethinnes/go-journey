package main

import "fmt"

func main() {
	ids := []int{33, 76, 54, 23, 11, 2}

	for i, id := range ids {
		//%d is base-ten integer, will be filled in by following arguments!
		fmt.Printf("%d - ID: %d\n", i, id)

	}
	//Without using index
	//for i, id := range ids {
	//fmt.Printf("ID: %d\n", i, id)

	//}
	//Add IDS together

	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	//Range with Map
	emails := map[string]string{"Luke": "luke@luke.com", "Jack": "jack@jack.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	//Just the key...

	for k := range emails {
		fmt.Println("Name: " + k)
	}
}
