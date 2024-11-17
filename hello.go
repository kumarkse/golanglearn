package main

import "fmt"

func main() {
	fmt.Println("hello world")
	fmt.Println("11" + "2w")
	fmt.Println(11 + 11)
	fmt.Println(11 + 11.11)
	fmt.Println(21.0 / 3)
	fmt.Println(true)
	fmt.Println(true && false)
	var ab = "55"
	fmt.Println(ab)
	var aba, bc = 3, 5
	fmt.Println(aba, bc)
	abc := "43.22"
	fmt.Println(abc)
	var name string = "abhi"
	fmt.Printf("my name is : %s ", name)
	//my name is : abhi

	name = "abhi"
	fmt.Printf("my name is : %T ", name)
	//my name is : string
	var smallvar float64 = 32.123456789101112131415
	fmt.Println(smallvar)
	trial1()

}
