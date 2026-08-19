# Go Concept Tutorial: Go Modules, Dependency Management, and Third-Party HTTP Routers

## 1. Overview and Core Concepts

Go Modules (`go.mod`) provide dependency management, module path resolution, and reproducible builds across Go projects.

Key principles covered in this concept:
- **Module Initialization**: `go mod init <module-name>` creates a `go.mod` file defining module dependencies and tracking explicit package versions.
- **Dependency Hygiene (`go mod tidy`)**: `go mod tidy` scans project imports, downloading missing third-party packages while removing unused dependencies from `go.mod` and `go.sum`.
- **Vendor Directory Isolation (`go mod vendor`)**: `go mod vendor` constructs a local `vendor/` directory containing complete source code copies of all external dependencies, ensuring offline reproducible builds.
- **Dependency Verification (`go mod verify`)**: `go mod verify` checks that local cached dependencies match expected cryptographic checksum hashes stored in `go.sum`.
- **Third-Party Routers**: Integrating third-party routers like `github.com/gorilla/mux` for robust HTTP request routing and pattern matching.

---

## 2. Code Walkthrough

```go
package main

import (
        "fmt"
        "log"
        "net/http"

        "github.com/gorilla/mux"
)

func main() {
        fmt.Println("Hello mod in golang")
        greeter()
        r := mux.NewRouter()
        r.HandleFunc("/", serveHome).Methods("GET")
        log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
        fmt.Println("Hey there mod users")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("<h1>Welcome to golang</h1>"))
}
```

### Explanation of Implementation
1. `r := mux.NewRouter()`: Initializes a Gorilla Mux router instance.
2. `r.HandleFunc("/", serveHome).Methods("GET")`: Registers an HTTP GET route at path `/` bound to the `serveHome` handler function.
3. `http.ListenAndServe(":4000", r)`: Binds the HTTP server to TCP port 4000, passing the Gorilla Mux router as the primary handler.

---

## 3. Toolchain and Execution

### Module Management Commands

Initialize module tracking:
```bash
go mod init mymodules
```

Fetch dependencies and clean `go.mod`:
```bash
go mod tidy
```

Create vendor folder for offline compilation:
```bash
go mod vendor
```

Verify dependency checksums:
```bash
go mod verify
```

### Running the Server
```bash
go run main.go
```

Test endpoint via `curl` or browser:
```bash
curl http://localhost:4000/
```

Expected response:
```html
<h1>Welcome to golang</h1>
```

---

## 4. Best Practices

- Always run `go mod tidy` prior to committing code to maintain clean dependency manifests.
- Commit `go.mod` and `go.sum` to source control to guarantee deterministic build behavior across build environments.
