package main

import (
	"bufio"
	"fmt"
	"os"
)

func trial1() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("give the number to be read : ")

	input, _ := reader.ReadString('\n')
	fmt.Println(input)
	fmt.Printf("type is %T", input)

}
