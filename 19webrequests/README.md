# Go Concept Tutorial: HTTP Web Requests and Stream Processing

## 1. Overview and Core Concepts

Making HTTP network calls in Go is facilitated by the standard `net/http` package.

Key principles covered in this concept:
- **HTTP Client GET Requests**: The `http.Get(url)` function issues an HTTP GET request to the designated URI, returning a pointer to an `http.Response` object.
- **Response Body Resource Management**: The `Response.Body` field is a readable `io.ReadCloser` stream. It is the strict caller's responsibility to close `res.Body` to prevent connection leaks and underlying TCP socket pool exhaustion.
- **Stream Reading**: Functions like `ioutil.ReadAll` (or `io.ReadAll` in modern Go) read raw bytes from an open stream until encountering an EOF (End-Of-File) signal.

---

## 2. Code Walkthrough

```go
package main

import (
        "fmt"
        "io/ioutil"
        "net/http"
)

const url = "https://example.com"

func main() {
        fmt.Println("LCO web requests")
        res, err := http.Get(url)
        if err != nil {
                panic(err)
        }
        fmt.Printf("Response is of type: %T\n", res)
        defer res.Body.Close() // Mandatory cleanup responsibility of caller

        databytes, err := ioutil.ReadAll(res.Body)
        if err != nil {
                panic(err)
        }
        content := string(databytes)
        fmt.Println(content)
}
```

### Explanation of Implementation
1. `res, err := http.Get(url)`: Sends an HTTP GET request to `https://example.com`.
2. `defer res.Body.Close()`: Schedules closure of the network response body stream when `main()` exits.
3. `databytes, err := ioutil.ReadAll(res.Body)`: Reads the entire HTML byte stream returned by the server.
4. `string(databytes)`: Converts raw response bytes to printable text.

---

## 3. Toolchain and Execution

Execute the source file:
```bash
go run main.go
```

Expected output snippet:
```text
LCO web requests
Response is of type: *http.Response
<!doctype html>
<html>
<head>
    <title>Example Domain</title>
...
```

---

## 4. Best Practices

- Always call `res.Body.Close()` inside a `defer` statement immediately after validating that `err == nil`.
- For production services, configure explicit client timeouts using `&http.Client{Timeout: 10 * time.Second}` rather than relying on `http.Get` which uses `http.DefaultClient` with no default timeout.
