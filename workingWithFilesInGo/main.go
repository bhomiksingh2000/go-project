package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("files")

	content := "my first go lang files example"
	file, err := os.Create("./myLogFile.txt")

	if err != nil {
		panic(err)
	} else {
		length, _ := io.WriteString(file, content)
		fmt.Printf("length of file is : %v\n", length)
		defer file.Close()
	}
	readFile()
}

func readFile() {
	databyte, err := os.ReadFile("./myLogFile.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(databyte))
}
