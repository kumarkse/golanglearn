package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	ip, _ := reader.ReadString('\n')
	// fmt.Print(ip)
	fmt.Printf("type : %T \n", ip)

	//newnum := ip +1  gives error
	newnum, err := strconv.ParseFloat(strings.TrimSpace(ip), 32)

	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Println(newnum + 23)
	}

	sttr := string(int(newnum))
	fmt.Printf("%T",sttr)
}
