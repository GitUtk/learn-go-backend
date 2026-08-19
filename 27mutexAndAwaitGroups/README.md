# Go Concept Tutorial: Race Conditions, Memory Safety, and Race Detection

## 1. Overview and Core Concepts

A data race (race condition) occurs when two or more goroutines access the same memory location concurrently, and at least one of the accesses is a write. Data races lead to non-deterministic bugs, corrupted state, and runtime crashes.

Key principles covered in this concept:
- **Race Conditions**: Occur when unsynchronized goroutines read and write shared data variables simultaneously (such as appending to a shared slice `score`).
- **Mutual Exclusion (`sync.Mutex`)**: Enforces critical section protection. Calling `mu.Lock()` blocks other goroutines from entering the critical section until `mu.Unlock()` is invoked.
- **Race Detector Tooling (`go run -race`)**: The Go compiler includes an integrated data race detector. Compiling or running code with the `-race` flag instruments memory access to identify unsynchronized memory operations at runtime.

---

## 2. Code Walkthrough

```go
package main

import (
        "fmt"
        "sync"
)

func main() {
        fmt.Println("Race condition - github.com")

        var wg sync.WaitGroup
        var mu sync.Mutex

        score := []int{0}

        wg.Add(3)

        go func() {
                defer wg.Done()
                fmt.Println("One R")

                mu.Lock()
                score = append(score, 1)
                mu.Unlock()
        }()

        go func() {
                defer wg.Done()
                fmt.Println("Two R")

                mu.Lock()
                score = append(score, 2)
                mu.Unlock()
        }()

        go func() {
                defer wg.Done()
                fmt.Println("Three R")

                mu.Lock()
                score = append(score, 3)
                mu.Unlock()
        }()

        wg.Wait()
        fmt.Println(score)
}
```

### Explanation of Implementation
1. `score := []int{0}`: Shared slice resource accessed by three separate concurrent goroutines.
2. `mu.Lock()` & `mu.Unlock()`: Each anonymous goroutine acquires `mu.Lock()` prior to calling `append(score, n)` and releases it immediately after. This guarantees that only one goroutine mutates `score` at any given millisecond.
3. `wg.Wait()`: Ensures `main()` waits for all three anonymous goroutines to finish before printing `score`.

---

## 3. Toolchain and Execution

### Running standard execution
```bash
go run main.go
```

Expected output (slice ordering depends on lock acquisition sequence):
```text
Race condition - github.com
One R
Two R
Three R
[0 1 2 3]
```

### Running with Race Detector
To verify that no race conditions exist:
```bash
go run -race main.go
```

If a data race were present (e.g. if `mu.Lock()` and `mu.Unlock()` were commented out), the race detector output would display detailed stack traces showing concurrent read/write locations:
```text
WARNING: DATA RACE
Write at 0x00c0000bc010 by goroutine 7:
...
```

---

## 4. Best Practices

- Always test concurrent Go applications using the `-race` flag (`go test -race ./...` or `go run -race main.go`).
- Keep critical sections protected by mutexes as short as possible to avoid bottlenecking concurrent goroutines.
- Consider `sync.RWMutex` when a shared resource experiences frequent read operations and infrequent write operations.
