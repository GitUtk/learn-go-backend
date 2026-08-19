# Go Concept Tutorial: Methods and Receiver Semantics

## 1. Overview and Core Concepts

A method in Go is a function bound to a specific receiver type. Methods allow defining behaviors on custom defined types and structures.

Key principles covered in this concept:
- **Method Receiver Syntax**: Method signatures place a receiver clause between the `func` keyword and the method identifier (`func (r ReceiverType) MethodName()`).
- **Value Receivers**: Defined as `(u User)`. The method receives a copy of the caller struct value. Modifications to fields inside value receiver methods operate on the copy and do not alter the caller's state.
- **Pointer Receivers**: Defined as `(u *User)`. The method receives a memory address pointer to the caller struct. Field modifications in pointer receiver methods alter the caller's actual memory state directly.

---

## 2. Code Walkthrough

```go
package main

import "fmt"

func main() {
        fmt.Println("Structs in golang")

        utkarsh := User{"Utkarsh", "test@test.com", true, 20}
        fmt.Println(utkarsh)
        fmt.Printf("%+v\n", utkarsh)
        fmt.Printf("Age is %v and email is %v\n", utkarsh.Age, utkarsh.Email)

        utkarsh.getStatus()
        utkarsh.NewMail()

        // Verifying that value receiver did not mutate caller state
        fmt.Println("Original user email after NewMail():", utkarsh.Email)
}

type User struct {
        Name   string
        Email  string
        Status bool
        Age    int
}

func (u User) getStatus() {
        fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
        u.Email = "test@go.dev"
        fmt.Println("Email of this user is: ", u.Email)
}
```

### Explanation of Implementation
1. `func (u User) getStatus()`: Value receiver method accessing `u.Status` without field mutation.
2. `func (u User) NewMail()`: Value receiver method modifying `u.Email`. Because `u` is passed by value (copied), setting `u.Email = "test@go.dev"` only affects the local copy inside `NewMail`. The caller's `utkarsh.Email` remains `"test@test.com"`.

---

## 3. Toolchain and Execution

Execute the source file:
```bash
go run main.go
```

Expected output:
```text
Structs in golang
{Utkarsh test@test.com true 20}
{Name:Utkarsh Email:test@test.com Status:true Age:20}
Age is 20 and email is test@test.com
Is user active:  true
Email of this user is:  test@go.dev
Original user email after NewMail(): test@test.com
```

---

## 4. Best Practices

- Use pointer receivers `(u *User)` whenever a method must modify the state of the receiver struct.
- Use pointer receivers for large structures to prevent copying overhead on every method call.
- Maintain consistency: avoid mixing value receivers and pointer receivers on the same struct type across a codebase.
