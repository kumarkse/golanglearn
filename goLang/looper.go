package main

import (
	"fmt"
	"math/rand"
)

func main() {

	randoms := rand.Perm(44)
	fmt.Print(randoms[20:28])

	nextrandoms := rand.NormFloat64()
	fmt.Println(nextrandoms)

	randoms = randoms[11:20]

	for i := 0; i < len(randoms); i++ {
		fmt.Print(randoms[i], " ")
	}

	fmt.Println()

	for i := range randoms {
		fmt.Print(randoms[i], " ")
	}

	fmt.Println()
	fmt.Println()
	fmt.Println()

	randomval := rand.Intn(len(randoms))
	fmt.Println("random number is : ", randomval)
	for indx, val := range randoms {
		if indx == randomval {
			fmt.Println("skipping ", indx, "th value ", val)
		}
		fmt.Printf("the value at index %v is %v\n", indx, val)

	}

	fmt.Println()
	fmt.Println()
	fmt.Println()
	i := 0
	for randomval > i {
		fmt.Println(i)
		i++

	}

	fmt.Println()
	fmt.Println()
	fmt.Println()

	for i := 1; randomval > i; {
		fmt.Println(i)
		i++
	}

	for i := 0; i < 10; i++ {
		if i == 4 {
			goto exit
		}
	}

exit:
	fmt.Print("exiting")

}
