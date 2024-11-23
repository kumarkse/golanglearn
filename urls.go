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
	// parse funnction thinks that the string is url and allows us to acces different components

	fmt.Println(request, err)

	fmt.Println(request.Scheme)
	request.Scheme = "changesScheme"
	fmt.Println(request)
	newUrl := request.String()
	fmt.Printf("converted string is %T ->%T :  %s",request,newUrl,newUrl)
	/***********************************************************************/

	
}
