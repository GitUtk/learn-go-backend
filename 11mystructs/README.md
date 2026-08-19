# Go Concept Tutorial: Custom Structs and Composition

## 1. Overview and Core Concepts

A `struct` (structure) in Go is a composite data type that groups together logical fields under a single user-defined type identifier. Structs provide object-oriented data aggregation without traditional class inheritance models.

Key principles covered in this concept:
- **Struct Type Definition**: Defined using the `type Name struct` syntax specifying named fields and data types.
- **No Class Inheritance**: Go omits traditional class hierarchies, `super` references, and parent-child inheritance in favor of explicit composition and interfaces.
- **Field Access**: Accessing and updating individual struct fields using dot notation (`instance.FieldName`).
- **Formatted Printing**: Formatting verbs in `fmt.Printf` enable detailed struct inspection: `%v` prints values, while `%+v` prints field names alongside their respective values.

---

## 2. Code Walkthrough

```go
package main

import "fmt"

func main() {
        fmt.Println("Structs in golang")
        // No inheritance in Go; no super or parent classes

        utkarsh := User{"Utkarsh", "test@test.com", true, 20}
        fmt.Println(utkarsh)
        fmt.Printf("%+v\n", utkarsh) // Detailed output including field names
        fmt.Printf("Age is %v and email is %v\n", utkarsh.Age, utkarsh.Email)
}

type User struct {
        Name   string
        Email  string
        Status bool
        Age    int
}
```

### Explanation of Implementation
1. `type User struct`: Declares a composite structure containing four fields (`Name`, `Email`, `Status`, `Age`).
2. `User{"Utkarsh", "test@test.com", true, 20}`: Instantiates a `User` struct using positional literal field mapping.
3. `fmt.Printf("%+v\n", utkarsh)`: Prints `{Name:Utkarsh Email:test@test.com Status:true Age:20}`.
4. `utkarsh.Age`: Direct access to the `Age` field value.

---

## 3. Toolchain and Execution

Execute the file:
```bash
go run main.go
```

Expected output:
```text
Structs in golang
{Utkarsh test@test.com true 20}
{Name:Utkarsh Email:test@test.com Status:true Age:20}
Age is 20 and email is test@test.com
```

---

## 4. Best Practices

- Prefer named struct field initialization (`User{Name: "Utkarsh", Age: 20}`) over positional literals to improve code readability and maintain compatibility when fields are added or reordered.
- Use exported field names (starting with capital letters) when fields must be marshaled to JSON or accessed outside the package.
