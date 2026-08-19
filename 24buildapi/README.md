# Go Concept Tutorial: In-Memory RESTful API Architecture with Gorilla Mux

## 1. Overview and Core Concepts

Building production-ready RESTful APIs in Go involves defining domain models, setting up route handlers, parsing HTTP request payloads, handling URL path variables, and returning structured JSON responses.

Key principles covered in this concept:
- **Domain Data Modeling**: Structuring data schemas (`Course`, `Author`) with JSON field tags for request/response serialization.
- **Route Mapping with Parameters**: Utilizing `mux.Vars(r)` to parse dynamic URL path variables (e.g., `/course/{id}`).
- **CRUD Operations**: Implementing Create, Read, Update, and Delete endpoints against an in-memory dataset slice.
- **Helper Validation Functions**: Writing internal validation methods on model structs (such as `IsEmpty()`) to validate inbound JSON payload content.
- **HTTP Status Codes and Error Handling**: Writing appropriate HTTP headers (`Content-Type: application/json`) and error status responses (`http.StatusBadRequest`, `http.StatusNotFound`).

---

## 2. Code Walkthrough

```go
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

// Model for courses
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

// Fake DB slice
var courses []Course

func (c *Course) IsEmpty() bool {
        return c.CourseName == ""
}

func main() {
        fmt.Println("Building RESTful API in Go")
        r := mux.NewRouter()

        // Seeding database
        courses = append(courses, Course{
                CourseId:    "2",
                CourseName:  "ReactJS",
                CoursePrice: 299,
                Author:      &Author{Fullname: "Utkarsh", Website: "go.dev"},
        })

        // Routing table
        r.HandleFunc("/", serveHome).Methods("GET")
        r.HandleFunc("/courses", getAllCourses).Methods("GET")
        r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
        r.HandleFunc("/course", createOneCourse).Methods("POST")
        r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
        r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

        log.Fatal(http.ListenAndServe(":4000", r))
}

func serveHome(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("<h1>Welcome to API by Go</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        params := mux.Vars(r)

        for _, course := range courses {
                if course.CourseId == params["id"] {
                        json.NewEncoder(w).Encode(course)
                        return
                }
        }
        json.NewEncoder(w).Encode("No Course found with given id")
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")

        if r.Body == nil {
                json.NewEncoder(w).Encode("Please send some data")
                return
        }

        var course Course
        _ = json.NewDecoder(r.Body).Decode(&course)
        if course.IsEmpty() {
                json.NewEncoder(w).Encode("No data inside JSON")
                return
        }

        rand.Seed(time.Now().UnixNano())
        course.CourseId = strconv.Itoa(rand.Intn(100))
        courses = append(courses, course)
        json.NewEncoder(w).Encode(course)
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        params := mux.Vars(r)

        for index, course := range courses {
                if course.CourseId == params["id"] {
                        courses = append(courses[:index], courses[index+1:]...)
                        var newCourse Course
                        _ = json.NewDecoder(r.Body).Decode(&newCourse)
                        newCourse.CourseId = params["id"]
                        courses = append(courses, newCourse)
                        json.NewEncoder(w).Encode(newCourse)
                        return
                }
        }
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        params := mux.Vars(r)

        for index, course := range courses {
                if course.CourseId == params["id"] {
                        courses = append(courses[:index], courses[index+1:]...)
                        json.NewEncoder(w).Encode("Record successfully deleted")
                        break
                }
        }
}
```

### Explanation of Implementation
1. `json.NewEncoder(w).Encode(courses)`: Directly serializes Go structures to JSON and streams them into the `http.ResponseWriter` output.
2. `params := mux.Vars(r)`: Extracts route path variables from the request context, such as `{id}` in `/course/{id}`.
3. `json.NewDecoder(r.Body).Decode(&course)`: Streams and deserializes the inbound HTTP request JSON body into a target struct.
4. `courses = append(courses[:index], courses[index+1:]...)`: Removes a target record by index during update and delete operations.

---

## 3. Toolchain and Execution

### Running the API Server
```bash
go run main.go
```

### Testing API Endpoints via cURL

Fetch all courses:
```bash
curl http://localhost:4000/courses
```

Fetch a single course by ID:
```bash
curl http://localhost:4000/course/2
```

Create a new course:
```bash
curl -X POST http://localhost:4000/course \
     -H "Content-Type: application/json" \
     -d '{"coursename":"Golang Masterclass","price":399,"author":{"fullname":"Utkarsh","website":"go.dev"}}'
```

Delete a course:
```bash
curl -X DELETE http://localhost:4000/course/2
```

---

## 4. Best Practices

- Use `json.NewDecoder(r.Body)` for processing HTTP body streams instead of reading all bytes into memory first with `ioutil.ReadAll`.
- Ensure appropriate HTTP headers (`Content-Type: application/json`) are set before writing responses to the response writer stream.
