# Go Concept Tutorial: HTTP Verbs, Request Payloads, and Form Encoding

## 1. Overview and Core Concepts

Building HTTP clients requires supporting different HTTP verbs (methods) such as `GET`, `POST` (JSON payloads), and `POST` (Form Data).

Key principles covered in this concept:
- **Sending JSON Payloads**: Using `strings.NewReader` to transform raw JSON string payloads into an `io.Reader` stream suitable for `http.Post`.
- **Form URL-Encoding**: Using `url.Values{}` to create key-value pairs encoded as `application/x-www-form-urlencoded` payloads via `http.PostForm`.
- **Efficient String Buffering**: Using `strings.Builder` to efficiently assemble response content from byte streams.

---

## 2. Code Walkthrough

```go
package main

import (
        "fmt"
        "io/ioutil"
        "net/http"
        "net/url"
        "strings"
)

func main() {
        fmt.Println("Welcome to web verb video - LCO")
        // PerformGetRequest()
        // PerformPostJsonRequest()
        defer PerformPostFormRequest()
}

func PerformGetRequest() {
        const myurl = "http://localhost:3000/get"
        res, err := http.Get(myurl)
        if err != nil {
                panic(err)
        }
        defer res.Body.Close()

        var responseString strings.Builder
        content, _ := ioutil.ReadAll(res.Body)
        byteCount, _ := responseString.Write(content)

        fmt.Println("Byte count:", byteCount)
        fmt.Println(responseString.String())
}

func PerformPostJsonRequest() {
        const myurl = "http://localhost:3000/post"

        requestBody := strings.NewReader(`
                {
                        "course":"Let's go with golang",
                        "price":0,
                        "platform":"github.com"
                }
        `)

        res, err := http.Post(myurl, "application/json", requestBody)
        if err != nil {
                panic(err)
        }
        defer res.Body.Close()

        content, _ := ioutil.ReadAll(res.Body)
        fmt.Println(string(content))
}

func PerformPostFormRequest() {
        const myurl = "http://localhost:3000/postform"

        data := url.Values{}
        data.Add("firstname", "utkarsh")
        data.Add("lastname", "singh")
        data.Add("email", "utkarsh@go.dev")

        res, err := http.PostForm(myurl, data)
        if err != nil {
                panic(err)
        }
        defer res.Body.Close()

        content, _ := ioutil.ReadAll(res.Body)
        fmt.Println(string(content))
}
```

### Explanation of Implementation
1. `strings.NewReader(...)`: Creates a reader supplying raw JSON bytes to `http.Post`.
2. `url.Values{}` & `data.Add(...)`: Constructs URL-encoded form data parameters (`firstname=utkarsh&lastname=singh...`).
3. `strings.Builder`: Provides a high-performance string accumulation buffer avoiding unnecessary memory allocations when processing response streams.

---

## 3. Toolchain and Execution

### Starting the Mock Server
This section depends on `server.js` running locally on port 3000. In a separate terminal run:
```bash
node server.js
```

### Running the Go Script
```bash
go run main.go
```

---

## 4. Best Practices

- Always specify the correct `Content-Type` header when sending request payloads (`application/json` for JSON, `application/x-www-form-urlencoded` for forms).
- Prefer `strings.Builder` over standard string concatenation (`+`) when accumulating string data inside loops or processing stream chunks.
