# Go Concept Tutorial: File System I/O Operations and Descriptor Cleanup

## 1. Overview and Core Concepts

File handling operations in Go are managed through the `os`, `io`, and `io/ioutil` standard library packages.

Key principles covered in this concept:
- **File Creation**: The `os.Create` function creates or truncates a file at a specified path, returning a writable `*os.File` descriptor pointer and an error object.
- **Stream Writing**: The `io.WriteString` utility writes string byte streams directly to an open `io.Writer` interface (such as a file descriptor).
- **File Reading**: `ioutil.ReadFile` (or `os.ReadFile` in modern Go) opens a target file, reads its total contents into a byte slice (`[]byte`), and closes the file descriptor automatically.
- **Resource Teardown**: Open file descriptors consume operating system resources. Using `defer file.Close()` guarantees descriptor teardown upon function completion.

---

## 2. Code Walkthrough

```go
package main

import (
        "fmt"
        "io"
        "io/ioutil"
        "os"
)

func main() {
        fmt.Println("Welcome to files in golang")
        content := "This needs to go in a file"

        file, err := os.Create("./mygofile.txt")
        checkNilErr(err)

        length, err := io.WriteString(file, content)
        checkNilErr(err)

        fmt.Println("Length is: ", length)
        readFile("mygofile.txt")
        defer file.Close()
}

func readFile(filename string) {
        databyte, err := ioutil.ReadFile(filename)
        checkNilErr(err)
        fmt.Println("Text data inside the file is\n", string(databyte))
}

func checkNilErr(err error) {
        if err != nil {
                panic(err)
        }
}
```

### Explanation of Implementation
1. `file, err := os.Create("./mygofile.txt")`: Opens or creates `mygofile.txt` with write permissions.
2. `length, err := io.WriteString(file, content)`: Writes the contents of string `content` to the file, returning the byte count written (26).
3. `readFile("mygofile.txt")`: Reads the raw byte array and converts it to a readable string format via `string(databyte)`.
4. `defer file.Close()`: Enforces file closure when `main()` exits.

---

## 3. Toolchain and Execution

Execute the source file:
```bash
go run main.go
```

Expected output:
```text
Welcome to files in golang
Length is:  26
Text data inside the file is
 This needs to go in a file
```

---

## 4. Best Practices

- Always defer closing file handles immediately after verifying that `os.Create` or `os.Open` succeeded without error.
- Use `os.ReadFile` and `os.WriteFile` (introduced in Go 1.16) as modern replacements for `ioutil.ReadFile` and `ioutil.WriteFile`.
- For large files, stream data using `bufio.Scanner` or `io.Copy` instead of loading entire file contents into memory at once.
