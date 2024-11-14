package main

import (
	"fmt"
	"sort"
)

func main() {
	/*fmt.Println("hello world")
	  fmt.Println("11"+"2w")
	  fmt.Println(11+11)
	  fmt.Println(11+11.11)
	  fmt.Println(21.0/3)
	  fmt.Println(true)
	  fmt.Println(true && false)
	  var ab="55"
	  fmt.Println(ab)
	  var ab  ,bc =3,5
	  fmt.Println(ab,bc)
	  ab:="43.22"
	  fmt.Println(ab)
	  var name string = "abhi"
	  fmt.Printf("my name is : %s ",name)
	  my name is : abhi

	  var name string = "abhi"
	  fmt.Printf("my name is : %T ",name)
	  my name is : string
	  var smallvar float64 = 32.123456789101112131415
	  fmt.Println(smallvar)
	  trial1()
	*/
	// array
	var list [5]int //compulsary number definintion

	list[0] = 5
	list[1] = 4
	list[2] = 3
	list[3] = 2
	list[4] = 1

	fmt.Println(list)
	fmt.Println(len(list))
	fmt.Printf("type :%T\n", list)
	// slice
	var list2 = []float32{22.0}
	fmt.Println(list2)

	list2 = append(list2, 22522, 256.2)
	fmt.Println(list2)

	list3 := append(list2[1:3])
	fmt.Println(list3)

	list4 := make([]int, 4)
	list4[0] = 11
	list4[1] = 14
	list4[2] = 13
	list4[3] = 12

	fmt.Println(list4)

	list4 = append(list4, 32, 33, 34) // will not give error now

	fmt.Println(list4)

	sort.Ints(list4)

	fmt.Println(list4)

	/*******************************************************************************************/
}
