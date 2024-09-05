package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"math/rand"

	"github.com/gorilla/mux"
)

// model for course and author [IN PRODUCTION THEY ARE IN SEPARATE FILES]
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"courseName"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

func main() {
	fmt.Println("Building API's in Go")
	r := mux.NewRouter()

	// filling in the data
	//seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS", CoursePrice: 299, Author: &Author{Fullname: "Hitesh Choudhary", Website: "lco.dev"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "MERN Stack", CoursePrice: 199, Author: &Author{Fullname: "Hitesh Choudhary", Website: "go.dev"}})

	// routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/courses/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/courses", createCourse).Methods("POST")
	r.HandleFunc("/courses/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/courses/{id}", deleteOneCourse).Methods("DELETE")

	// listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}

// fake data
var courses []Course

// middleware - helper in Go
func (c *Course) IsEmpty() bool {

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

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "applicatioan/json")

	// grab id from request
	params := mux.Vars(r)

	// loop through courses, find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with given id")

}

func createCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("adding a course in our DB")
	w.Header().Set("Content-Type", "applicatioan/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("The Request is Empty")
	}

	var course Course

	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	// generate unique id, string
	// append course into courses

	seed := time.Now().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	course.CourseId = strconv.Itoa(rng.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Updating one course")
	w.Header().Set("Content-Type", "applicatioan/json")

	params := mux.Vars(r)
	idFound := false

	for index, course := range courses {
		if course.CourseId == params["id"] {
			idFound = true
			// now first we will delete the index from the courses,
			courses = append(courses[:index], courses[index+1:]...)

			// now update the course in req
			var courseInReq Course
			_ = json.NewDecoder(r.Body).Decode(&courseInReq)
			courseInReq.CourseId = params["id"]
			courses = append(courses, courseInReq)
			json.NewEncoder(w).Encode(courseInReq)

			return
		}
	}

	//TODO: send a response when id is not found
	if !idFound {
		json.NewEncoder(w).Encode("No data inside JSON")

	}

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting this particular course from our fake DB Courses")
	w.Header().Set("Content-Type", "applicatioan/json")

	params := mux.Vars(r)
	idFound := false
	for index, course := range courses {
		if course.CourseId == params["id"] {
			fmt.Println("Deleting this course with id ", params["id"])
			courses = append(courses[:index], courses[index+1:]...)
			idFound = true
			break
		}
	}

	if !idFound {
		json.NewEncoder(w).Encode("No data inside JSON")

	}

}
