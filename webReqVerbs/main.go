package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// GetRequest()
	// postRequest()
	PostFormRequest()
}

func GetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)

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

func postRequest() {
	// for our post API , only JSON data can be used (as per the API contract)

	const myurl = "http://localhost:8000/post"

	requestBody := strings.NewReader(`

		{

			"coursename":"Let's go with golang",

			"price": 0,

			"platform":"learnCodeOnline.in"

		}

	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println("content is  ", string(content))

}

/*
Form Submissions: Contact forms, registration forms,
and feedback forms typically use POST to send data to the server.
File Uploads: Uploading images, documents, or other files to a server
is done via a POST request.
*/
func PostFormRequest() {
	const urlStr = "http://localhost:8000/postform"

	// Create form data
	data := url.Values{}
	data.Add("firstname", "hitesh")
	data.Add("lastname", "choudhary")
	data.Add("email", "hitesh@go.dev")

	// Send POST request
	response, err := http.PostForm(urlStr, data)

	if err != nil {
		panic(err)
	}

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))

}
