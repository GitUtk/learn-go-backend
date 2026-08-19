# Go Concept Tutorial: Type Conversion and String Sanitation

## 1. Overview and Core Concepts

Input gathered from console streams or external files arrives as text data (`string`). Performing mathematical calculations or structural operations on user input requires converting raw string formats into explicit numerical data types.

Key principles covered in this concept:
- **String Sanitation**: User input read via buffered streams retains trailing delimiter bytes (such as `\n` or `\r\n`). Functions like `strings.TrimSpace` purge whitespace and trailing line terminators.
- **Type Parsing**: The `strconv` package converts basic types to and from string representations. Specifically, `strconv.ParseFloat` parses text strings into floating-point numbers of specified precision bit sizes.
- **Idiomatic Error Checking**: Evaluating returned `error` values explicitly using `if err != nil` control blocks to maintain execution safety.

---

## 2. Code Walkthrough

```go
package main

import (
        "bufio"
        "fmt"
        "os"
        "strconv"
        "strings"
)

func main() {
        fmt.Println("Welcome to our pizza app")
        fmt.Println("Please rate out pizza between 1 and 5")

        reader := bufio.NewReader(os.Stdin)
        input, _ := reader.ReadString('\n')

        fmt.Println("Thanks for rating, ", input)

        numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

        if err != nil {
                fmt.Println(err)
        } else {
                fmt.Println("Added 1 to your rating: ", numRating+1)
        }
}
```

### Explanation of Implementation
1. `strings.TrimSpace(input)`: Strips leading and trailing whitespace, spaces, and newline characters from the user-entered input string.
2. `strconv.ParseFloat(..., 64)`: Parses the cleaned string into a `float64` floating-point value. The second parameter specifies a 64-bit precision result.
3. `if err != nil`: Checks whether parsing succeeded or failed. If invalid characters (e.g., non-numeric text) were passed, `err` holds details of the parsing error; otherwise `err` is `nil`, allowing safe arithmetic execution (`numRating + 1`).

---

## 3. Toolchain and Execution

Run the script interactively:
```bash
go run main.go
```

Sample interactive session with valid numeric input:
```text
Welcome to our pizza app
Please rate out pizza between 1 and 5
4.5
Thanks for rating,  4.5

Added 1 to your rating:  5.5
```

Sample session with invalid non-numeric input:
```text
Welcome to our pizza app
Please rate out pizza between 1 and 5
five
Thanks for rating,  five

strconv.ParseFloat: parsing "five": invalid syntax
```

---

## 4. Best Practices

- Never attempt numerical type casting or arithmetic directly on raw user inputs without sanitizing whitespace first.
- Always check returned errors from string conversion utilities before relying on converted numerical variables.
