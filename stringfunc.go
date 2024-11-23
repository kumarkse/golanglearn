package main

import (
	"fmt"
	"strings"
)

func main() {
	data1 := "apple,banana,orange"
	parts := strings.Split(data1, ",")
	fmt.Println(parts, len(parts))

	data2 := "appleABbananaABorange"
	parts = strings.Split(data2, "AB")
	fmt.Println(parts, len(parts))

	data3 := "   loremipsumsdhisghdbchd   bckhdgidka    "
	count := strings.Count(data3, "d")
	fmt.Println(count)

	trimmed:= strings.TrimSpace(data3)
	fmt.Println(trimmed)

	data4 := trimmed + data2 
	fmt.Println(data4)

	splittedjoined := strings.Join(parts,"-")
	fmt.Println(splittedjoined)
}
