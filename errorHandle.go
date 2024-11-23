package main

import "fmt"

func main() {
	fmt.Println("error handling module")
	val , _ := try(54, 0)
	fmt.Println(val)

}

func try(a, b int) (int, error) {
	if b == 0{
		return 0, fmt.Errorf("Divide by zero error")
	}
	return a / b, nil
}
