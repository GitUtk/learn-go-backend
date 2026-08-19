# Go Concept Tutorial: Pointers and Memory Address Manipulation

## 1. Overview and Core Concepts

A pointer is a variable that stores the memory address of another variable rather than storing a direct data value. Pointers allow programs to pass references to memory locations across functions and modify variables in place without making memory copies.

Key principles covered in this concept:
- **Memory Address Operator (`&`)**: The address-of operator yields the underlying memory location address of a variable (e.g., `&myNumber`).
- **Dereference Operator (`*`)**: The dereference operator accesses or alters the data value held at the memory address pointed to by the pointer variable (e.g., `*ptr`).
- **Pointer Types**: Declared using an asterisk preceding the type name (e.g., `*int` represents a pointer to an integer).
- **Zero Value of Pointers**: Uninitialized pointer variables default to `nil`.

---

## 2. Code Walkthrough

```go
package main

import "fmt"

func main() {
        fmt.Println("Welcome to class on pointers")

        // Uninitialized pointer declaration yields nil
        // var ptr *int
        // fmt.Println("Value of pointer is ", ptr)

        myNumber := 23
        var ptr = &myNumber

        fmt.Println("Value of actual pointer is ", ptr)  // Prints hex memory address
        fmt.Println("Value of actual pointer is ", *ptr) // Prints value stored at memory address

        *ptr = *ptr + 2
        fmt.Println("New value is ", myNumber) // Underlying variable modified directly
}
```

### Explanation of Implementation
1. `var ptr = &myNumber`: Creates a pointer variable `ptr` holding the physical memory address of `myNumber`.
2. `fmt.Println(ptr)`: Outputs the hexadecimal memory location address (e.g., `0xc000012088`).
3. `fmt.Println(*ptr)`: Dereferences `ptr`, retrieving the value `23`.
4. `*ptr = *ptr + 2`: Mutates the underlying memory contents directly. Consequently, accessing `myNumber` afterwards yields `25`.

---

## 3. Toolchain and Execution

Execute the file:
```bash
go run main.go
```

Sample output:
```text
Welcome to class on pointers
Value of actual pointer is  0xc000012088
Value of actual pointer is  23
New value is  25
```

---

## 4. Best Practices

- Use pointers when passing large structures to functions to avoid costly memory copying operations.
- Use pointer receivers when methods need to mutate receiver struct states.
- Always check for `nil` pointers prior to dereferencing to prevent runtime panic crashes.
