package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// EncodingJson()
	decodeJson()
}

// type course struct {
// 	Name     string
// 	Price    int
// 	Link     string
// 	Password string
// 	Tags     []string
// }

// betterment : means whenever the Course object will be converted to json
// then CourseId will be printed / logged as courseid

type course struct {
	Name     string `json : "courseName"`
	Price    int
	Link     string   `json : "website"`
	Password string   `json : "-"`
	Tags     []string `json : "tags, omitEmpty"`
}

func EncodingJson() {
	SliceOfCourses := []course{
		{"React.js", 112, "abc.in", "344234", []string{"web-dev", "frontend"}},
		{"Node.js", 113, "qwd.in", "332432", []string{"web-dev", "backend"}},
		{"Go", 1134, "svav.in", "43141341", []string{"web-dev", "backend"}},
	}

	// finalData, err := json.Marshal(SliceOfCourses)
	/*
		The json.MarshalIndent function in Go is used to convert a Go data structure
		(like structs, slices, maps, etc.) into a JSON format with indentation for readability.
		This is particularly useful when you want to produce human-readable JSON,
		such as when you're logging,
		 debugging, or returning formatted JSON from an API.
	*/
	finalData, err := json.MarshalIndent(SliceOfCourses, "", "\t")

	// but the problem here is that

	if err != nil {
		panic(err)
	}
	fmt.Println(string(finalData))
}

func decodeJson() {

	// becuase web se data humesha bytes mai aata h
	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJS Bootcamp",
		"Price": 299,
		"website": "LearnCodeOnline.in",
		"tags": ["web-dev","js"]
	}
	`)

	var convertedCourseFromJson course

	checkIfJsonIsValid := json.Valid(jsonDataFromWeb)

	/*
		func Unmarshal(data []byte, v interface{}) error

			The key point here is that json.Unmarshal requires a pointer
			to the variable where the parsed JSON data should be stored.
			This is because the function needs to modify the contents of
			that variable, and in Go, you can only modify a variable outside
			of a function if you have a pointer to it.
	*/

	if checkIfJsonIsValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &convertedCourseFromJson)
		fmt.Printf("%#v\n", convertedCourseFromJson)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	/*
		Without the &, you would be passing a copy of the convertedCourseFromJson variable,
		and json.Unmarshal would modify the copy, not the original variable.
		After the function call, the original variable would remain unchanged,
		which is not what you want.
	*/

	// some cases where you just want to add data to key value
	// no need of creating a struct here as it is the json
	// will be converted to the required object

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for key, value := range myOnlineData {
		fmt.Printf("key is : %v , value is : %v, and type is : %T", key, value, value)
	}

}
