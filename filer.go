package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("entereed te filer fn")

	content := "to be put into the file second time"

	file, err := os.Create("./contentor.txt")

	if err != nil {
		fmt.Println("we got error and returning")
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		fmt.Println("we got error and returning")
		panic(err)
	}

	fmt.Println(length)

	file.Close()

	/*********************************************/

	filepath := "contentor.txt"

	filenew, err := os.Open(filepath)
	buffer := make([]byte, 1024)
	bytesread, err := filenew.Read(buffer)
	fmt.Println("Read values : ", string(buffer[:bytesread]))

	/*********************************************/

	databyte, err := ioutil.ReadFile(filepath)

	fmt.Println("Read values again : ", string(databyte))

	/*********************************************/

	databyte2, err := os.ReadFile(filepath)

	fmt.Println("Read values once again : ", string(databyte2))

}
