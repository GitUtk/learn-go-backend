# Go Concept Tutorial: Functions, Multiple Return Signatures, and Variadic Parameters

## 1. Overview and Core Concepts

Functions in Go are first-class constructs defined using the `func` keyword. They accept typed parameters and can return zero, one, or multiple values.

Key principles covered in this concept:
- **Function Signatures**: Explicitly state input parameter types and output return types (`func name(param Type) ReturnType`).
- **No Named Nested Functions**: Standard named function declarations cannot be nested inside other functions in Go. (Anonymous functions and closures, however, can be assigned to local variables within function bodies).
- **Multiple Return Signatures**: Functions in Go can return multiple values simultaneously, commonly used for returning both a result and an `error` object (`(int, string)`).
- **Variadic Functions**: Functions that accept a variable number of final parameters of a specific type using the ellipsis operator (`values ...int`). Inside the function body, variadic parameters are received as a slice (`[]int`).

---

## 2. Code Walkthrough

```go
package main

import "fmt"

func main() {
        fmt.Println("Welcome to functions in go lang")
        greeter()

        greeterTwo()
        result := adder(3, 5)
        fmt.Println("Result is: ", result)

        sum, message := proAdder(345, 345, 345, 34, 5345, 4)
        fmt.Println("My message is ", message, "and sum is ", sum)
}

func greeter() {
        fmt.Println("Hello from golang")
}

func greeterTwo() {
        fmt.Println("Another method")
}

func adder(a int, b int) int {
        return a + b
}

func proAdder(values ...int) (int, string) {
        total := 0
        for _, value := range values {
                total += value
        }
        return total, "It works"
}
```

### Explanation of Implementation
1. `adder(a int, b int) int`: Expects two `int` parameters and returns a single `int`.
2. `proAdder(values ...int) (int, string)`: Takes any number of `int` arguments, iterates through them via a `range` loop inside a slice context, and returns both an integer sum and a confirmation string tuple.
3. `sum, message := proAdder(...)`: Receives the tuple returned by `proAdder`.

---

## 3. Toolchain and Execution

Execute the source file:
```bash
go run main.go
```

Expected output:
```text
Welcome to functions in go lang
Hello from golang
Another method
Result is:  8
My message is  It works and sum is  6418
```

---

## 4. Best Practices

- Use multiple return values to return explicit error objects alongside operation results (`(ResultType, error)`).
- Use variadic parameters when writing utility routines that accept arbitrary numbers of input arguments (such as formatting or logging helpers).
