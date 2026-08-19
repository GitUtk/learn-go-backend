# Go Concept Tutorial: Variables, Data Types, and Visibility

## 1. Overview and Core Concepts

Go is a statically typed language, meaning that variable types are evaluated and verified at compile time. Variable declarations specify the variable name, data type, and optionally an initial memory value.

Key principles covered in this concept:
- **Explicit Variable Typing**: Declaring a variable with an explicit type (`var identifier type = value`).
- **Default Zero Values**: Variables declared without explicit initial values default to their type's zero value (`0` for integers, `0.0` for floats, `""` for strings, and `false` for booleans).
- **Implicit Type Inference**: Omitting explicit type annotations when initializing with `var`, allowing the Go compiler to infer the data type from the assignment expression.
- **Short Declaration Operator (`:=`)**: Block-scoped variable initialization shorthand available strictly within function bodies.
- **Constant Declarations (`const`)**: Immutable value bindings fixed at compile time.
- **Exported vs. Unexported Identifiers (Visibility Scope)**: Capitalization rules dictate package scope. Identifiers beginning with a capital letter (e.g., `LoginToken`) are exported (public), while lowercase identifiers are unexported (private to the package).

---

## 2. Code Walkthrough

```go
package main

import "fmt"

const LoginToken string = "kjaDFasdabgsg" // Public (Exported)

func main() {
        var username string = "Utkarsh"
        fmt.Println(username)
        fmt.Printf("Variable is of type: %T \n", username)

        var isLoggedIn bool = true
        fmt.Println(isLoggedIn)
        fmt.Printf("Varible is of type : %T \n", isLoggedIn)

        var smallVal int = 255
        fmt.Println(smallVal)
        fmt.Printf("Varible is of type : %T \n", smallVal)

        var smallFloat float64 = 255.4567457674
        fmt.Println(smallFloat)
        fmt.Printf("Varible is of type : %T \n", smallFloat)

        var anotherVariable int
        fmt.Println(anotherVariable)
        fmt.Printf("Varible is of type : %T \n", anotherVariable)

        // Implicit type inference
        var website = "github.com"
        fmt.Println(website)

        // Short declaration operator
        numberOfUser := 300000.0
        fmt.Println(numberOfUser)

        fmt.Println(LoginToken)
        fmt.Printf("Varible is of type : %T \n", LoginToken)
}
```

### Explanation of Implementation
1. `const LoginToken string = "kjaDFasdabgsg"`: Package-level exported constant. Accessible outside the current package due to uppercase naming convention.
2. `var username string = "Utkarsh"`: Explicit declaration of string type.
3. `fmt.Printf("%T \n", variable)`: Uses the `%T` verb in `fmt.Printf` to print the concrete data type of the specified variable.
4. `var anotherVariable int`: Uninitialized declaration. Printing yields the integer zero value `0`.
5. `var website = "github.com"`: Type inferred as `string`. Assigning a different type (such as an integer) post-declaration causes a compilation error.
6. `numberOfUser := 300000.0`: Short declaration operator inferring float type (`float64`).

---

## 3. Toolchain and Execution

Execute the source file directly:
```bash
go run main.go
```

Expected output:
```text
Utkarsh
Variable is of type: string 
true
Varible is of type : bool 
255
Varible is of type : int 
255.4567457674
Varible is of type : float64 
0
Varible is of type : int 
github.com
300000
kjaDFasdabgsg
Varible is of type : string 
```

---

## 4. Best Practices

- Use the short declaration operator `:=` inside function bodies to reduce redundant code, reserving explicit `var` declarations for zero-value initialization or package-level variables.
- Follow Go's export visibility conventions: use PascalCase for exported entities and camelCase for private/package-internal entities.
- Choose specific floating-point precision (`float32` vs `float64`) based on application precision requirements.
