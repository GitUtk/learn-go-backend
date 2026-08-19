# Go Concept Tutorial: Deferred Stack Execution and Resource Cleanup

## 1. Overview and Core Concepts

The `defer` statement postpones the execution of a function call until the surrounding function returns.

Key principles covered in this concept:
- **Last-In-First-Out (LIFO) Execution Order**: Multiple deferred statements inside a function are pushed onto an execution stack. When the surrounding function exits, deferred statements execute in reverse order (LIFO).
- **Delayed Execution**: Deferred calls execute immediately before the surrounding function terminates (either via normal return or runtime panic).
- **Immediate Parameter Evaluation**: Arguments passed to a deferred function are evaluated immediately when the `defer` line is reached, not when the deferred function actually executes.
- **Resource Management**: `defer` is predominantly used for guaranteed resource teardown (e.g., closing file descriptors, unlocking mutexes, closing network connections).

---

## 2. Code Walkthrough

```go
package main

import "fmt"

func main() {
        defer fmt.Println("World") // Stack position 3 (Executes last)
        defer fmt.Println("One")   // Stack position 2
        defer fmt.Println("Two")   // Stack position 1 (Executes first among main defers)

        fmt.Println("Hello")

        myDefer()
}

func myDefer() {
        for i := 0; i < 5; i++ {
                defer fmt.Println(i) // Pushes 0, 1, 2, 3, 4 onto stack
        }
}
```

### Explanation of Implementation
1. In `main()`, three statements are deferred: `"World"`, `"One"`, and `"Two"`.
2. `fmt.Println("Hello")` executes immediately.
3. `myDefer()` is invoked. Inside `myDefer()`, a loop defers printing `0`, `1`, `2`, `3`, and `4`. When `myDefer()` finishes, its deferred stack pops and prints `4 3 2 1 0`.
4. Finally, as `main()` exits, its deferred stack pops in LIFO order, printing `"Two"`, `"One"`, and `"World"`.

---

## 3. Toolchain and Execution

Execute the source file:
```bash
go run main.go
```

Expected output:
```text
Hello
4
3
2
1
0
Two
One
World
```

---

## 4. Best Practices

- Always place `defer file.Close()` or `defer mu.Unlock()` immediately after successfully acquiring a resource or error check.
- Be cautious when deferring inside long-running loops; deferred statements accumulate on the call stack until the surrounding function exits, which can delay memory release or cause descriptor exhaustion.
