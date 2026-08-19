# Go Concept Tutorial: Goroutines, WaitGroups, and Concurrent HTTP Polling

## 1. Overview and Core Concepts

Concurrency is a fundamental design feature of Go. Goroutines are lightweight threads managed directly by the Go runtime scheduler rather than the operating system.

Key principles covered in this concept:
- **Goroutine Spawning**: Prefixing a function call with the `go` keyword launches execution asynchronously in a new goroutine (e.g., `go getStatusCode(web)`).
- **WaitGroups (`sync.WaitGroup`)**: A WaitGroup waits for a collection of goroutines to finish executing.
  - `wg.Add(delta)` increments the WaitGroup counter by delta.
  - `wg.Done()` decrements the counter by 1 (typically invoked via `defer wg.Done()` inside the worker goroutine).
  - `wg.Wait()` blocks execution of the main thread until the WaitGroup counter drops to zero.
- **Mutex Locks (`sync.Mutex`)**: A mutual exclusion lock prevents data race conditions when multiple concurrent goroutines attempt to modify shared memory variables simultaneously.

---

## 2. Code Walkthrough

```go
package main

import (
        "fmt"
        "net/http"
        "sync"
)

var signals = []string{"test"}

var wg sync.WaitGroup // WaitGroup instance pointer
var mut sync.Mutex    // Mutex instance pointer

func main() {
        websiteList := []string{
                "https://github.com",
                "https://go.dev",
                "https://vercel.com",
                "https://youtube.com",
                "https://google.com",
        }

        for _, web := range websiteList {
                go getStatusCode(web)
                wg.Add(1)
        }
        wg.Wait() // Blocks until all website status checks finish
        fmt.Println(signals)
}

func getStatusCode(endpoint string) {
        defer wg.Done()

        res, err := http.Get(endpoint)
        if err != nil {
                fmt.Println("OOPS in endpoint")
        } else {
                mut.Lock()
                signals = append(signals, endpoint)
                mut.Unlock()
                fmt.Printf("%d stats code for %s\n", res.StatusCode, endpoint)
        }
}
```

### Explanation of Implementation
1. `go getStatusCode(web)`: Spawns a concurrent goroutine to perform an HTTP GET request for each URL in `websiteList`.
2. `wg.Add(1)` & `defer wg.Done()`: Increments the counter before spawning each thread and decrements it when `getStatusCode` exits.
3. `mut.Lock()` & `mut.Unlock()`: Guards the shared `signals` slice update. Without `mut.Lock()`, appending to `signals` concurrently from multiple goroutines causes a data race.

---

## 3. Toolchain and Execution

Execute the source file:
```bash
go run main.go
```

Sample output (order of completion varies due to concurrent execution):
```text
200 stats code for https://go.dev
200 stats code for https://github.com
200 stats code for https://google.com
200 stats code for https://vercel.com
200 stats code for https://youtube.com
[test https://go.dev https://github.com https://google.com https://vercel.com https://youtube.com]
```

---

## 4. Best Practices

- Always pass `sync.WaitGroup` pointers to functions (or access them via package scope), as passing a WaitGroup by value copies its internal state and causes deadlocks.
- Ensure every `wg.Add(N)` call has a corresponding `wg.Done()` call, ideally placed inside `defer wg.Done()` to prevent premature or blocked calls to `wg.Wait()`.
- Use mutexes to protect all shared state mutations occurring across concurrent goroutines.
