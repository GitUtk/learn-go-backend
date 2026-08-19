# Go Concept Tutorial: Static Arrays and Memory Allocation

## 1. Overview and Core Concepts

An array in Go is a fixed-size contiguous sequence of elements of a single specified data type. The length of an array is part of its type definition (`[4]string` and `[3]string` are distinct, incompatible types).

Key principles covered in this concept:
- **Fixed Sizing**: Array lengths are fixed at compile time and cannot be resized dynamically.
- **Value Semantics**: Arrays in Go are value types, not reference types. Assigning an array to another variable or passing it to a function creates a complete copy of all array elements.
- **Default Element Initialization**: Elements of an array initialized without values default to their type's zero value (e.g., empty strings `""` for string arrays).
- **Array Literals**: Instantiating and initializing arrays directly using brace syntax `[N]Type{...}`.

---

## 2. Code Walkthrough

```go
package main

import "fmt"

func main() {
        fmt.Println("Welcome to array in golang")

        var fruitList [4]string
        fruitList[0] = "apple"
        fruitList[1] = "mango"
        fruitList[3] = "Peach"

        fmt.Println("Fruit list is: ", fruitList)
        fmt.Println("Fruit list is:", len(fruitList)) // Evaluates to 4

        var vegList = [3]string{"potato", "beans", "mushroom"}

        fmt.Println("Vegy list is: ", vegList)
}
```

### Explanation of Implementation
1. `var fruitList [4]string`: Allocates memory for 4 string elements. `fruitList[2]` remains unassigned and holds an empty string `""`.
2. `len(fruitList)`: Returns the total capacity length allocated to the array (4), regardless of how many non-zero elements have been assigned.
3. `vegList := [3]string{...}`: Declares and populates a 3-element string array literal directly.

---

## 3. Toolchain and Execution

Execute the source file:
```bash
go run main.go
```

Expected output:
```text
Welcome to array in golang
Fruit list is:  [apple mango  Peach]
Fruit list is: 4
Vegy list is:  [potato beans mushroom]
```

---

## 4. Best Practices

- Use arrays primarily when element count is fixed and known at compile time (such as matrix transformations, fixed hash digest buffers, or low-level memory buffers).
- For dynamic data collections, use slices (`[]T`) rather than static arrays.
