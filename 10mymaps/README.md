# Go Concept Tutorial: Hash Maps, Deletion, and Map Iteration

## 1. Overview and Core Concepts

A map in Go is a built-in unordered hash table collection mapping unique keys of a specified type to values of another type (`map[KeyType]ValueType`).

Key principles covered in this concept:
- **Map Initialization**: Uninitialized map variables default to `nil`. Attempting to insert keys into a `nil` map triggers a runtime panic. Maps must be initialized using `make(map[K]V)` or map literals.
- **Key Deletion**: The `delete(mapInstance, key)` built-in function removes specified keys from a map. If the target key does not exist, `delete` executes as a silent no-op without error.
- **Iteration Order**: Iterating over maps using `for key, value := range mapInstance` produces key-value pairs in randomized, non-deterministic order across program runs.

---

## 2. Code Walkthrough

```go
package main

import "fmt"

func main() {
        fmt.Println("Maps in golang")

        language := make(map[string]string)
        language["JS"] = "Javascript"
        language["RB"] = "Ruby"
        language["PY"] = "Python"

        fmt.Println("List of all languages: ", language)
        fmt.Println("JS shorts for: ", language["JS"])

        delete(language, "RB")
        fmt.Println("List of all languages: ", language)

        // Iterating over key-value map entries
        for key, value := range language {
                fmt.Printf("For key %v, value is %v\n", key, value)
        }
}
```

### Explanation of Implementation
1. `language := make(map[string]string)`: Allocates an active, empty map instance with `string` keys and `string` values.
2. `language["JS"] = "Javascript"`: Inserts key `"JS"` mapped to value `"Javascript"`.
3. `delete(language, "RB")`: Removes the key `"RB"` and its associated value from the map.
4. `for key, value := range language`: Iterates through remaining map entries, assigning key and value identifiers during each step.

---

## 3. Toolchain and Execution

Execute the source file:
```bash
go run main.go
```

Expected output:
```text
Maps in golang
List of all languages:  map[JS:Javascript PY:Python RB:Ruby]
JS shorts for:  Javascript
List of all languages:  map[JS:Javascript PY:Python]
For key JS, value is Javascript
For key PY, value is Python
```

---

## 4. Best Practices

- Always initialize maps using `make()` before attempting key-value assignments.
- Use the comma-ok idiom (`value, exists := mapInstance[key]`) to distinguish between non-existent map keys and keys stored with zero values.
- Note that Go maps are not concurrency-safe for concurrent read-write operations; synchronization primitives (`sync.RWMutex` or `sync.Map`) are required in multithreaded environments.
