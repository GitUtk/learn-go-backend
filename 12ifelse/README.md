# Go Concept Tutorial: Conditional Branching and Inline Statements

## 1. Overview and Core Concepts

Conditional execution in Go is structured around `if`, `else if`, and `else` statements.

Key principles covered in this concept:
- **Boolean Expressions**: Conditions must evaluate explicitly to a boolean value (`true` or `false`). Go does not support implicit conversion of integers or pointers to booleans (no "truthy" or "falsy" values).
- **Inline Variable Initialization**: Go permits declaring and initializing a variable directly inside the conditional statement header (`if initialization; condition { ... }`). Variables declared this way are scoped exclusively to the `if`, `else if`, and `else` blocks.
- **Parentheses Omission**: Go idiomatic style omits surrounding parentheses around conditional boolean expressions while requiring enclosing curly braces `{}`.

---

## 2. Code Walkthrough

```go
package main

import "fmt"

func main() {
        fmt.Println("If/else in golang")
        loginCount := 4
        var result string

        if loginCount < 10 {
                result = "Regular User"
        } else if loginCount > 10 {
                result = "Watch out"
        } else {
                result = "Excatly 10 login count"
        }
        fmt.Println(result)

        if 9%2 == 0 {
                fmt.Println("Number is even")
        } else {
                fmt.Println("Number is odd")
        }

        // Inline statement initialization within condition scope
        if num := 3; num < 10 {
                fmt.Println("Num is less than 10")
        } else {
                fmt.Println("Num is not less than 10")
        }
}
```

### Explanation of Implementation
1. `if loginCount < 10`: Standard evaluation of integer comparison.
2. `9%2 == 0`: Uses modulo arithmetic operator `%` to test numerical parity.
3. `if num := 3; num < 10`: Initializes variable `num` with value `3` and immediately evaluates `num < 10`. The variable `num` is out of scope once the `if/else` block completes.

---

## 3. Toolchain and Execution

Execute the source file:
```bash
go run main.go
```

Expected output:
```text
If/else in golang
Regular User
Number is odd
Num is less than 10
```

---

## 4. Best Practices

- Keep conditional branches simple and eliminate unnecessary `else` statements by returning early (guard clauses).
- Use inline variable initializations when a temporary variable is needed solely for the conditional evaluation scope.
