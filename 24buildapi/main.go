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

// middleware helper-file
func (c *Course) isEmpty() bool {
	//return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("api with go")

	r := mux.NewRouter()

	//seeding
	courses = append(courses, Course{CourseId: "21", CourseName: "ReactJs", CoursePrice: 499, Author: &Author{Fullname: "Hitesh Choudary", Website: "lco.dev"}})
	courses = append(courses, Course{CourseId: "28", CourseName: "NodeJs", CoursePrice: 699, Author: &Author{Fullname: "Hitesh Choudary", Website: "lco.dev"}})
	courses = append(courses, Course{CourseId: "71", CourseName: "GO", CoursePrice: 899, Author: &Author{Fullname: "Hitesh Choudary", Website: "lco.dev"}})

	//routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	//listen to port
	log.Fatal(http.ListenAndServe(":4000", r))
}

//controller -file

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API on GO lang</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get one course")
	w.Header().Set("Content-Type", "application/json")

	//grab id from request
	params := mux.Vars(r)

	//loop through courses and find matching id and return response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create one course")
	w.Header().Set("Content-Type", "application/json")

	//what if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("please sent some data")
		return
	}

	//if data is {}
	var newCourse Course
	_ = json.NewDecoder(r.Body).Decode(&newCourse)
	if newCourse.isEmpty() {
		json.NewEncoder(w).Encode("please sent some data")
		return
	}

	//if course already exist
	for _, course := range courses {
		if course.CourseName == newCourse.CourseName {
			json.NewEncoder(w).Encode("Course already exist")
			return
		}
	}

	//generate unique id in string
	//append new course to courses
	rand.Seed(time.Now().UnixNano())
	newCourse.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, newCourse)
	json.NewEncoder(w).Encode(newCourse)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create one course")
	w.Header().Set("Content-Type", "application/json")

	//grab id from request
	params := mux.Vars(r)

	//loop, id, remove, add with my ID
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with this ID")
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create one course")
	w.Header().Set("Content-Type", "application/json")

	//grab id from request
	params := mux.Vars(r)

	//loop, id, remove, add with my ID
	for index, course := range courses {
		if course.CourseId == params["id"] {
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with this ID")
}
