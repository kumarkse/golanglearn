package main

import (
	"fmt"
	"io"
	"net/http"
)

const urll = "https://www.iitgoa.ac.in/~nehak/example.html"

func main() {
	fmt.Println("accessing web request")
	response, err := http.Get(urll)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%T", response)
	defer response.Body.Close() // callers responsibility

	content,err:=io.ReadAll(response.Body)

	content2 := string(content)

	fmt.Println(content2)

	
	fmt.Printf("%v",response.StatusCode)




}
