# Go Concept Tutorial: Unified Iteration, Range Loops, and Control Jumps

## 1. Overview and Core Concepts

Go unifies all looping mechanisms under a single keyword: `for`. Go does not feature `while` or `do-while` keywords; instead, different configurations of `for` achieve counter loops, condition loops, infinite loops, and range iterations.

Key principles covered in this concept:
- **Condition-Only Loops**: A `for` loop with only a condition expression functions identically to a traditional `while` loop (`for condition { ... }`).
- **Range Iteration**: The `for index, value := range collection` construct iterates over elements of slices, arrays, maps, strings, or channels.
- **Loop Control Keywords**: `break` terminates loop execution immediately, while `continue` skips the remainder of the current iteration.
- **Labeled Control Statements (`goto`)**: The `goto` statement jumps execution unconditionally to a defined statement label within the same function.

---

## 2. Code Walkthrough

```go
package main

import "fmt"

func main() {
        fmt.Println("Welcome to loops in golang")

        days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}
        fmt.Println(days)

        // Standard index-based loop:
        // for i := 0; i < len(days); i++ { fmt.Println(days[i]) }

        // Range index loop:
        // for i := range days { fmt.Println(days[i]) }

        // Range index-value loop:
        // for index, day := range days { fmt.Printf("index is %v and value is %v\n", index, day) }

        rougueValue := 1

        for rougueValue < 10 {
                if rougueValue == 5 {
                        goto lco
                }
                fmt.Println("Value is: ", rougueValue)
                rougueValue++
        }

lco:
        fmt.Println("Jumping at github.com")
}
```

### Explanation of Implementation
1. `days := []string{...}`: Slice containing day names.
2. `for rougueValue < 10`: Executes continuously as long as `rougueValue` is less than `10`.
3. `if rougueValue == 5 { goto lco }`: When `rougueValue` reaches 5, execution branches unconditionally to the label `lco:`, bypassing the rest of the loop.
4. `lco:`: Code execution target label.

---

## 3. Toolchain and Execution

Execute the source file:
```bash
go run main.go
```

Expected output:
```text
Welcome to loops in golang
[Sunday Tuesday Wednesday Friday Saturday]
Value is:  1
Value is:  2
Value is:  3
Value is:  4
Jumping at github.com
```

---

## 4. Best Practices

- Use `for range` loops when working with slices, arrays, maps, and channels for clean and idiomatic iteration.
- Use the blank identifier `_` to ignore unused index or value parameters during `range` loops (e.g., `for _, day := range days`).
- Minimize the use of `goto` statements, as uncontrolled branching can reduce code readability and maintainability.
