package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://abcd.com"

func main() {
	fmt.Println("lets start Url management")
	fmt.Println(myurl)

	request, err := url.Parse(myurl)

	fmt.Println(request, err)

	fmt.Print(request.Port())

	/***********************************************************************/

	
}
