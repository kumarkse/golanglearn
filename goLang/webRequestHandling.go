package main

import (
	"fmt"
	"io"
	"net/http"
)

const urll = "https://www.iitgoa.ac.in/~nehak/example.html"
const urls = "https://jsonplaceholder.typicode.com/todos/1"

func main() {
	fmt.Println("accessing web request")
	response, err := http.Get(urls)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%T", response)
	defer response.Body.Close() // callers responsibility

	content,err:=io.ReadAll(response.Body)

	content2 := string(content)

	fmt.Println(content2)

	
	fmt.Printf("%v",response.StatusCode)


	/********************************************/

}
