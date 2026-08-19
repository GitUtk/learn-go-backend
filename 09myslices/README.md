# Go Concept Tutorial: Dynamic Slices, Allocation, and Element Removal

## 1. Overview and Core Concepts

Slices are dynamic abstraction wrappers built on top of underlying contiguous Go arrays. A slice header consists of three components: a pointer to the backing array, a length (`len`), and a capacity (`cap`).

Key principles covered in this concept:
- **Dynamic Resizing**: Slices grow dynamically using the built-in `append` function. When appending elements exceeds current slice capacity, Go automatically allocates a larger backing array, copies existing elements, and updates the pointer.
- **Slice Header Allocation (`make`)**: The `make([]T, length, capacity)` function allocates zeroed backing arrays with specified initial lengths and optional capacities.
- **Sub-slicing Syntax**: Slices can be re-sliced using range indexing syntax `slice[low:high]`.
- **Sorting Utilities**: Standard library package `sort` provides functions like `sort.Ints` and `sort.IntsAreSorted` for sorting slice contents.
- **Element Removal Pattern**: Removing an element at a given index is achieved by variadically appending the remaining slice segments: `append(slice[:index], slice[index+1:]...)`.

---

## 2. Code Walkthrough

```go
package main

import (
        "fmt"
        "sort"
)

func main() {
        fmt.Println("Introduction to slicing on the array")
        var fruitList = []string{"Apple", "Tomato", "Peach"}
        fmt.Printf("Type of fruiteList is %T\n", fruitList)

        fruitList = append(fruitList, "Mango", "Banana")
        fmt.Println(fruitList)

        fruitList = append(fruitList[1:3])
        fmt.Println(fruitList)

        highScores := make([]int, 4)
        highScores[0] = 234
        highScores[1] = 945
        highScores[2] = 465
        highScores[3] = 867

        highScores = append(highScores, 555, 666, 777) // Memory reallocation

        fmt.Println(highScores)

        sort.Ints(highScores) // Sorts slice elements in ascending order
        fmt.Println(highScores)
        fmt.Println(sort.IntsAreSorted(highScores)) // Evaluates to true

        // Removing element at index 2 ("go") from slice
        var courses = []string{"python", "java", "go", "cpp", "reactjs"}
        fmt.Println(courses)
        var index int = 2
        courses = append(courses[:index], courses[index+1:]...)
        fmt.Println(courses)
}
```

### Explanation of Implementation
1. `append(fruitList, "Mango", "Banana")`: Appends elements dynamically to the slice.
2. `fruitList[1:3]`: Creates a sub-slice starting at index 1 up to (but excluding) index 3.
3. `make([]int, 4)`: Instantiates a slice of integers with length 4 initialized to zero values `[0, 0, 0, 0]`.
4. `sort.Ints(highScores)`: Sorts integer slice values in-place in ascending order.
5. `courses = append(courses[:index], courses[index+1:]...)`: Removes element at index 2 by combining sub-slice `[0, 1]` with variadic elements of sub-slice `[3, 4]`.

---

## 3. Toolchain and Execution

Execute the file:
```bash
go run main.go
```

Expected output:
```text
Introduction to slicing on the array
Type of fruiteList is []string
[Apple Tomato Peach Mango Banana]
[Tomato Peach]
[234 945 465 867 555 666 777]
[234 465 555 666 777 867 945]
true
[python java go cpp reactjs]
[python java cpp reactjs]
```

---

## 4. Best Practices

- Pre-allocate slice capacity with `make([]T, 0, expectedCapacity)` when the number of elements is known beforehand to eliminate redundant memory reallocations during `append`.
- Remember that sub-slicing (`s[1:3]`) shares the same underlying array memory. Modifying elements in a sub-slice alters the parent slice.
