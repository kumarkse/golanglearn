package main

import (
	"fmt"
	"io"
	"net/http"
)
type Message struct {
    Name string
    Body string
    Time int64
}
func main() {
	
	urll:="http://127.0.0.1:5000/gettut"
	data, _ := http.Get(urll)
	defer data.Body.Close()
	content, _ := io.ReadAll(data.Body)


	fmt.Println(data.ContentLength)
	fmt.Println(string(content))
	fmt.Printf("%T",content)
}