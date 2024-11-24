package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	// making request GET
	const geturl = "http://127.0.0.1:5000"
	const puturl = "http://127.0.0.1:5000/posttut"

	// getRequest(geturl)
	// getRequest("http://127.0.0.1:5000/gettut")

	postRequest(puturl)

}

func getRequest(url string) {
	data, _ := http.Get(url)
	defer data.Body.Close()
	content, _ := io.ReadAll(data.Body)
	fmt.Println(data.ContentLength)
	fmt.Println(string(content))

	/* other way*/

	var stringBuilderWay strings.Builder
	// hhtp.get -> io.readall -> stringbuilder
	abc, _ := stringBuilderWay.Write(content)
	content2 := stringBuilderWay.String()
	fmt.Println(abc)
	fmt.Println(content2)
}

func postRequest(url string) {

	contentType := "application/json"

	// self creation of json
	contentToSend := strings.NewReader(`
		{
			"value" : 502,
			"message" : "helllow"
		}
	`)

	resp, _ := http.Post(url, contentType, contentToSend)

	defer resp.Body.Close()

	respRead, _ := io.ReadAll(resp.Body)
	fmt.Println(string(respRead))

}
