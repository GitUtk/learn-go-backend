# Go Concept Tutorial: Channel Communication Primitives and Goroutine Synchronization

## 1. Overview and Core Concepts

Channels in Go are typed conduit pipelines through which concurrent goroutines communicate and synchronize execution by sending and receiving values (`ch <- value` and `value := <-ch`). Channels embody Go's core concurrency philosophy: *"Do not communicate by sharing memory; instead, share memory by communicating."*

Key principles covered in this concept:
- **Unbuffered Channels**: Initialized using `make(chan Type)`. Unbuffered channels have no internal storage capacity. Send operations (`ch <- v`) block until another goroutine executes a corresponding receive operation (`<-ch`), making unbuffered channels synchronous primitives.
- **Buffered Channels**: Initialized with a capacity parameter (`make(chan Type, capacity)`). Send operations block only when the channel buffer is full, and receive operations block only when the buffer is empty.
- **Channel Operations**:
  - Send value: `ch <- value`
  - Receive value: `val := <-ch`
  - Close channel: `close(ch)`
- **Directional Channels**: Restricting channels inside function parameters to send-only (`chan<- T`) or receive-only (`<-chan T`) for type safety.

---

## 2. Code Walkthrough

```go
package main

import (
        "fmt"
        "sync"
)

func main() {
        fmt.Println("Channels in golang - github.com/gitUtk")

        myCh := make(chan int)
        wg := &sync.WaitGroup{}

        wg.Add(2)

        // Sender goroutine
        go func(ch chan int, wg *sync.WaitGroup) {
                defer wg.Done()

                fmt.Println("Sending 5...")
                ch <- 5
                fmt.Println("Sent 5")
        }(myCh, wg)

        // Receiver goroutine
        go func(ch chan int, wg *sync.WaitGroup) {
                defer wg.Done()

                value := <-ch
                fmt.Println("Received:", value)
        }(myCh, wg)

        wg.Wait()

        fmt.Println("Main function completed")
}
```

### Explanation of Implementation
1. `myCh := make(chan int)`: Creates an unbuffered channel for transferring `int` values between goroutines.
2. `ch <- 5` (Sender): Attempts to send `5` onto `myCh`. The sender blocks until the receiver goroutine is ready.
3. `value := <-ch` (Receiver): Receives the value `5` from `myCh`, unblocking the sender goroutine.
4. `wg.Wait()`: Blocks until both sender and receiver goroutines call `wg.Done()`.

---

## 3. Toolchain and Execution

Execute the source file:
```bash
go run main.go
```

Expected output:
```text
Channels in golang - github.com/gitUtk
Sending 5...
Sent 5
Received: 5
Main function completed
```

---

## 4. Best Practices

- Always ensure a receiving goroutine is actively listening before sending data to an unbuffered channel to prevent permanent goroutine deadlocks (`fatal error: all goroutines are asleep - deadlock!`).
- Only the sending goroutine should close a channel (`close(ch)`). Sending data to a closed channel causes a runtime panic.
- Use channel directionality (`chan<- int` or `<-chan int`) in function parameter definitions to enforce strict compile-time communication boundaries.
