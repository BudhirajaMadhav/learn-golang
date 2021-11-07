package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for course - file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// middleware, helper - file
func (c *Course) IsEmpty() bool {
	
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""


}

func main() {

	fmt.Println("API - by Madhav Budhiraja")
	r := mux.NewRouter()

	// seeding 
	courses = append(courses, Course{CourseId: "2", CourseName: "reactjs", CoursePrice: 858, Author: &Author{Fullname: "Max Millian", Website: "max.mil"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "mern-stack", CoursePrice: 745, Author: &Author{Fullname: "Achut jain", Website: "ac.hut"}})

	// routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")


	// listen to port
	log.Fatal(http.ListenAndServe(":4000", r))

}

// controlers - file

// serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to this API</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Get all courses")
	w.Header().Set("content-type","application/json")
	json.NewEncoder(w).Encode(courses)

}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	
	fmt.Println("Get one course")	
	w.Header().Set("content-type","application/json")
	
	// grab id from request
	params := mux.Vars(r)

	// loop through the courses, find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode(string("No course found with id of"+ params["id"]))

}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	
	fmt.Println("Create new course")
	w.Header().Set("content-type", "application/json")

	// what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	// what about  {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}
	
	// TODO(): check only if title is duplicate

	// generate a new id
	// append this new course into courses'

	rand.Seed(time.Now().Unix())
	course.CourseId = strconv.Itoa(rand.Intn(100)) 
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)

}

func updateOneCourse(w http.ResponseWriter, r *http.Request)  {
	
	fmt.Println("Update one course")
	w.Header().Set("content-type", "application/json")

	// grab id from request
	params := mux.Vars(r)

	// loop, id, remove, add with myid

	for index, course := range courses {
		if course.CourseId == params["id"] {
		
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		
		}
	}

	// TODO(): send a response when id is not found

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("content-type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index + 1:]...)
			json.NewEncoder(w).Encode(string("Course with courseID" + params["id"] + "Deleted"))
			return
		}
	}

	json.NewEncoder(w).Encode("This course doesn't exist")

}