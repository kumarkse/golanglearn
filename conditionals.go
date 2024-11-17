package main

import (
	"fmt"
	"math/rand"
)

func main() {
	ifelsetrial := []int{5, 2, 5, 1515, 4, 4, 5, 45, 4184, 588, 485, 48, 48, 545, 48, 48, 54, 5}

	for i := 0; i < len(ifelsetrial); i++ {

		if ifelsetrial[i]%2 == 0 {
			fmt.Printf("the value at index %v is even , i.e. %v\n", i, ifelsetrial[i])
		} else if ifelsetrial[i]%3 == 0 {

			fmt.Println("divisible by 3")
		} else {
			fmt.Println("odd value")

		}

		if k := 4; ifelsetrial[i]%k == 0 {
			fmt.Println("divisible by 4")
		}

	}

	switch rand.Intn(7) {
	case 1:
		fmt.Println(1, " is printed")
	case 2:
		fmt.Println(2, " is printed")
	case 3:
		fmt.Println(3, " is printed")
	case 4:
		fmt.Println(4, " is printed")
	case 5:
		fmt.Println(5, " is printed")
	case 6:
		fmt.Println(6, " is printed")
	default:
		fmt.Println("nothing ")
	}
}
