package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// model for course and author [IN PRODUCTION THEY ARE IN SEPARATE FILES]
type Course struct {
	CourseId    string  `json : courseid`
	CourseName  string  `json : courseName`
	CoursePrice int     `json : price`
	Author      *Author `json : author`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

func main() {
	fmt.Println("Building API's in Go")
}

// fake data
var courses []Course

// middleware - helper in Go
func (c *Course) isEmpty() bool {

	/*
		In Go, when you define a method with a pointer receiver
		(e.g., func (c *Course) IsEmpty() bool), it means that the
		\ method is associated with the Course type and can modify
		the actual Course instance it’s called on, since it has access
		 to the memory address of that instance.
	*/
	return c.CourseId == "" && c.CourseName == ""
}

// controllers [IN PRODUCTION THEY ARE IN SEPARATE FILES]

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to homepage</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all courses")
	w.Header().Set("Content-Type", "applicatioan/json")
	json.NewEncoder(w).Encode(courses)
}
