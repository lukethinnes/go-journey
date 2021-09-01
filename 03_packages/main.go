package main

import (
	"fmt"
	"math"

	"github.com/lukethinnes/go-journey/03_packages/strutil"
)

func main() {
	fmt.Println(math.Floor(2.5))
	fmt.Println(math.Ceil(2.8))
	fmt.Println(math.Sqrt(25))
	fmt.Println(strutil.Reverse("olleh"))
}
