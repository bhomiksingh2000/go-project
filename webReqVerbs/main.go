package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	GetRequest()
}

func GetRequest() {
	const url = "http://localhost:8000/get"

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)

	// doesn't actually print the content of the body;
	// it just prints the reference to the io.ReadCloser object.
	// To print the actual content of the body, you'd need to read it,

	// fmt.Println("Body is ", response.Body)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)

	// METHOD - 2
	byteCount, _ := responseString.Write(content)

	fmt.Println("ByteCount is :", byteCount)
	fmt.Println("Acutal content is :", responseString.String())

	// METHOD - 1
	// fmt.Println(string(content))

	// fmt.Println(content) // in bytes

	/*
		When to Prefer Each Method:
		Method 1 is preferred for simple cases where you
		just want to print or handle the content directly
		as a string without further modifications.


		Method 2 is preferred when you are working with
		large amounts of data or need to build a string incrementally,
		 as it offers better performance and flexibility.


		In your example, Method 1 is likely more suitable unless
		you specifically need the benefits of strings.Builder,
		such as when dealing with multiple pieces of data or
		performing string operations in a loop.


	*/
}
