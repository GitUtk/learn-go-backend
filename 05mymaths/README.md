# Go Concept Tutorial: Mathematics, Big Integers, and Random Number Generation

## 1. Overview and Core Concepts

Go provides multiple standard packages for performing arithmetic operations, generating random numbers, and computing arbitrary-precision mathematics.

Key principles covered in this concept:
- **Pseudo-Random vs Cryptographic Randomness**: The `math/rand` package implements pseudo-random number generators suitable for casual non-security tasks. In contrast, `crypto/rand` provides cryptographically secure random number generation backed by operating system entropy streams.
- **Arbitrary-Precision Arithmetic**: The `math/big` package supplies structures (such as `big.Int`, `big.Float`, `big.Rat`) for working with numerical values that exceed standard fixed-width integer capacities (e.g., 64-bit limits).
- **Entropy Streams**: Functions in `crypto/rand` accept `rand.Reader` as a source of system entropy.

---

## 2. Code Walkthrough

```go
package main

import (
        "crypto/rand"
        "fmt"
        "math/big"
)

func main() {
        fmt.Println("Welcome to maths in golang")

        // Cryptographically secure random integer generation within range [0, 5)
        myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
        fmt.Println(myRandomNum)
}
```

### Explanation of Implementation
1. `big.NewInt(5)`: Constructs a new `big.Int` object representing the upper non-inclusive bound limit (5).
2. `rand.Int(rand.Reader, big.NewInt(5))`: Generates a uniform random integer in the range `[0, max)` using system entropy from `crypto/rand.Reader`. The returned value is of type `*big.Int`.

---

## 3. Toolchain and Execution

Execute the source file:
```bash
go run main.go
```

Expected output (value varies per execution):
```text
Welcome to maths in golang
3
```

---

## 4. Best Practices

- For security-sensitive applications (such as key generation, token generation, or cryptography), always use `crypto/rand` instead of `math/rand`.
- When using `math/rand` in modern Go versions (1.20+), prefer automatic seeding over manual calls to `rand.Seed()`.
- Explicitly convert data types when performing arithmetic between mismatched scalar types (e.g., converting `float64` to `int` before addition).
