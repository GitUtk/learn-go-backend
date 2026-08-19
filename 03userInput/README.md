# Go Concept Tutorial: User Input Processing and Buffered Stream Reading

## 1. Overview and Core Concepts

Handling standard user input in command-line Go applications requires managing input stream buffers. The Go standard library provides the `bufio` and `os` packages to facilitate efficient I/O operations over system input streams (`os.Stdin`).

Key principles covered in this concept:
- **Buffered Reader Stream**: The `bufio.NewReader` function wraps `os.Stdin` in a buffer, preventing memory overhead from frequent system calls during stream reads.
- **Delimiter Reading**: The `ReadString('\n')` method reads input sequentially from the buffer until encountering a delimiter character (e.g., the newline character `\n`).
- **Multiple Return Signatures and Error Idioms**: Methods in Go frequently return tuple signatures `(result, error)`.
- **Blank Identifier (`_`)**: The blank identifier `_` acts as a write-only value placeholder used to discard unwanted return values or ignore errors explicitly during prototype exploration.

---

## 2. Code Walkthrough

```go
package main

import (
        "bufio"
        "fmt"
        "os"
)

func main() {
        welcome := "Welcome to user input"
        fmt.Println(welcome)
        reader := bufio.NewReader(os.Stdin) // For taking input we use bufio
        fmt.Println("Enter the rating for our pizza:")

        // comma ok  || err ok syntax
        input, _ := reader.ReadString('\n')
        fmt.Println("Thanks for rating, ", input)
        fmt.Printf("Type of this rating is %T : ", input)
}
```

### Explanation of Implementation
1. `reader := bufio.NewReader(os.Stdin)`: Constructs a buffered reader tied to standard system input stream (`os.Stdin`).
2. `input, _ := reader.ReadString('\n')`: Invokes `ReadString` with `\n` as the termination character. The function returns both the captured string (including the newline delimiter) and an error object. The error object is explicitly ignored using `_`.
3. `fmt.Printf("Type of this rating is %T : ", input)`: Demonstrates that input read from `bufio.Reader` is returned as type `string`.

---

## 3. Toolchain and Execution

Execute the file interactively:
```bash
go run main.go
```

Sample interactive session:
```text
Welcome to user input
Enter the rating for our pizza:
4
Thanks for rating,  4

Type of this rating is string : 
```

---

## 4. Best Practices

- While ignoring errors via `_` simplifies quick scripts, production applications must handle non-nil errors returned by `ReadString`.
- Note that `ReadString('\n')` retains the trailing newline delimiter in the returned string buffer; input sanitation (such as `strings.TrimSpace`) is required prior to numerical conversions.