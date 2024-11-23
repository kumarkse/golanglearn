package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

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
	var list2 = []float32{22.0, 456}
	fmt.Println(list2)

	list2 = append(list2, 22522, 256.2)
	fmt.Println(list2)

	list3 := list2[1:3]
	fmt.Println("->", list3)

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

	var course = []string{"abc", "cde", "fgl", "hhf", "fddf"}

	reader := bufio.NewReader(os.Stdin)

	removeidx, _ := reader.ReadString('\n')
	removeidx2, _ := strconv.ParseInt(strings.TrimSpace(removeidx), 10, 64)
	course2 := append(course[:removeidx2], course[(removeidx2+1):]...)

	fmt.Println(course2)
	/*************************************************************************/

	ds1 := make(map[int]string)
	ds1[1] = "js"
	ds1[2] = "dj"
	ds1[3] = "dshdl"
	fmt.Println(ds1)

	fmt.Println(ds1[1], ds1[2])
	delete(ds1, 1)
	fmt.Println(ds1)
	value,exists := ds1[2]

	fmt.Println(value,exists)


	ds2 := map[int]string {
		1:"bob",
		2:"bobby",
		3:"cappy",
	}

	fmt.Println(ds2)

	reader2 := bufio.NewReader(os.Stdin)
	scanner, _ := reader2.ReadString('\n')
	// str := "11"
	intstr, _ := strconv.Atoi(strings.TrimSpace(scanner))
	fmt.Println(intstr)

	/**********************************************************************************/

	valu:=90
	var ptr * int= &valu
	fmt.Println(ptr)

	pttr := &valu

	fmt.Println(pttr)


}
