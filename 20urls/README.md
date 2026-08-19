# Go Concept Tutorial: URL Parsing, Query Parameter Extraction, and Construction

## 1. Overview and Core Concepts

Uniform Resource Identifiers (URIs) and URLs are parsed, processed, and constructed in Go using the standard `net/url` package.

Key principles covered in this concept:
- **URL Decomposition**: The `url.Parse(rawURL)` function parses a raw URL string into a structured `*url.URL` instance exposing component fields: `Scheme`, `Host`, `Path`, `Port()`, and `RawQuery`.
- **Query Parameter Maps**: Calling `.Query()` on a `*url.URL` object parses URL encoded query parameters into a `url.Values` map representation (`map[string][]string`).
- **URL Reconstruction**: Constructing dynamic URLs cleanly by initializing a `&url.URL{...}` struct instance and calling its `.String()` method.

---

## 2. Code Walkthrough

```go
package main

import (
        "fmt"
        "net/url"
)

const myurl string = "https://something.com:3000/learn?course=btech&class=it"

func main() {
        fmt.Println("Handling URLs in goLang")
        fmt.Println(myurl)
        result, _ := url.Parse(myurl)

        fmt.Println("Scheme:", result.Scheme)
        fmt.Println("Host:", result.Host)
        fmt.Println("RawQuery:", result.RawQuery)
        fmt.Println("Path:", result.Path)
        fmt.Println("Port:", result.Port())

        qparams := result.Query()
        fmt.Printf("The type of query params are : %T\n", qparams)

        fmt.Println("Class query value:", qparams["class"])
        for _, val := range qparams {
                fmt.Println("Params is ", val)
        }

        // Constructing a new URL programmatically
        partsOfUrl := &url.URL{
                Scheme:  "https",
                Host:    "localhost",
                Path:    "/user",
                RawPath: "user=admin",
        }

        anotherUrl := partsOfUrl.String()
        fmt.Println("Constructed URL:", anotherUrl)
}
```

### Explanation of Implementation
1. `url.Parse(myurl)`: Decomposes `https://something.com:3000/learn?course=btech&class=it`.
2. `result.Port()`: Extracts port number `"3000"`.
3. `result.Query()`: Converts query string `course=btech&class=it` into a map where `qparams["class"]` evaluates to `[it]`.
4. `partsOfUrl.String()`: Rebuilds string representation `https://localhost/user`.

---

## 3. Toolchain and Execution

Execute the source file:
```bash
go run main.go
```

Expected output:
```text
Handling URLs in goLang
https://something.com:3000/learn?course=btech&class=it
Scheme: https
Host: something.com:3000
RawQuery: course=btech&class=it
Path: /learn
Port: 3000
The type of query params are : url.Values
Class query value: [it]
Params is  [btech]
Params is  [it]
Constructed URL: https://localhost/user
```

---

## 4. Best Practices

- Always use `url.Parse` and `url.Values` for manipulating query parameters rather than manual string concatenation to ensure proper URL escaping.
- Remember that `url.Values` maps keys to slices of strings (`[]string`), accounting for query parameters that appear multiple times in a single URL (e.g. `?tag=go&tag=web`).
