# Go Programming Curriculum and Reference Manual

Welcome to the Go Programming Reference codebase. This repository contains a structured, end-to-end curriculum designed to take engineers from fundamental programming concepts in Go to production-grade backend development, HTTP API engineering, MongoDB persistence, and concurrent execution patterns.

---

## Directory Index

### Module 1: Go Language Fundamentals
Basic syntax, environment initialization, type inference, standard input/output streams, numerical parsing, mathematical packages, and temporal management.

| Section | Concept Tutorial | Description | Primary Package / Feature |
| :--- | :--- | :--- | :--- |
| 01 | [Hello World & Environment](01hello/README.md) | Entry point definition, standard package output, and Go module initialization. | `main`, `fmt`, `go mod` |
| 02 | [Variables & Data Types](02variables/README.md) | Static typing, implicit type inference, short declarations, public/private naming scope. | `var`, `const`, `:=` |
| 03 | [User Input Processing](03userInput/README.md) | Buffered stream reading, system standard input, and blank identifier error ignoring. | `bufio.Reader`, `os.Stdin` |
| 04 | [Type Conversion](04conversion/README.md) | String parsing, floating-point parsing, whitespace trimming, and error evaluation. | `strconv`, `strings.TrimSpace` |
| 05 | [Mathematics & Randomness](05mymaths/README.md) | Cryptographically secure pseudo-random generation and arbitrary-precision integers. | `crypto/rand`, `math/big` |
| 06 | [Time & Binary Compilation](06mytime/README.md) | Temporal representation, reference layout formatting, and cross-platform compilation. | `time.Now`, `GOOS` build flags |

---

### Module 2: Memory & Data Structures
Pointers, direct memory modification, static contiguous arrays, dynamic slices, key-value hash maps, and custom composite structures.

| Section | Concept Tutorial | Description | Primary Package / Feature |
| :--- | :--- | :--- | :--- |
| 07 | [Pointers & Memory Access](07mypointers/README.md) | Address-of operator, pointer dereferencing, and direct memory state manipulation. | `&`, `*` operators |
| 08 | [Static Arrays](08myarray/README.md) | Fixed-size array declaration, memory layout, array literals, and length evaluation. | `[N]T` array types |
| 09 | [Dynamic Slices](09myslices/README.md) | Slice headers, dynamic reallocation with append, sub-slicing syntax, and element removal. | `[]T`, `make`, `sort` |
| 10 | [Hash Maps](10mymaps/README.md) | Hash map initialization, key insertion, key deletion, and map iteration. | `map[K]V`, `delete`, `range` |
| 11 | [Structs & Type Schemas](11mystructs/README.md) | Custom composite type definition, field assignment, and type composition models. | `type T struct` |

---

### Module 3: Control Structures & Code Organization
Branching statements, expression switching, unified iteration semantics, function signatures, receiver methods, and deferred stack execution.

| Section | Concept Tutorial | Description | Primary Package / Feature |
| :--- | :--- | :--- | :--- |
| 12 | [Conditional Branching](12ifelse/README.md) | Boolean expression evaluation, inline conditional initialization statements, and branching. | `if`, `else if`, `else` |
| 13 | [Switch Statements](13switchcase/README.md) | Value matching, default fallbacks, implicit break execution, and fallthrough semantics. | `switch`, `case`, `fallthrough` |
| 14 | [Iteration & Loops](14loops/README.md) | Unified for loop patterns, slice range iteration, control break/continue, and labeled jumps. | `for`, `range`, `goto` |
| 15 | [Functions & Variadic Arguments](15functions/README.md) | Function definitions, multiple return signatures, and variadic parameter handling. | `func`, `...int` |
| 16 | [Methods & Receiver Functions](16methods/README.md) | Struct method binding, value receiver semantics, and receiver modification limits. | `func (r Receiver)` |
| 17 | [Deferred Execution Stack](17defer/README.md) | LIFO execution stack, delayed evaluation of parameters, and resource cleanup mechanics. | `defer` keyword |

---

### Module 4: File System I/O, Web Requests & Data Serialisation
File descriptor creation, byte stream reading, HTTP client interaction, URL decomposition, HTTP verb simulation, and JSON field tagging.

| Section | Concept Tutorial | Description | Primary Package / Feature |
| :--- | :--- | :--- | :--- |
| 18 | [File I/O Operations](18files/README.md) | File creation, string stream writing, whole-file byte reading, and descriptor closing. | `os`, `io`, `io/ioutil` |
| 19 | [HTTP Web Requests](19webrequests/README.md) | HTTP client GET operations, response body reading, and stream closure responsibility. | `net/http`, `io.ReadAll` |
| 20 | [URL Parsing & Decomposition](20urls/README.md) | URI parsing, scheme/host/path extraction, query parameter maps, and URL reconstruction. | `net/url` |
| 21 | [HTTP Verbs & Payload Forms](21webreqverbs/README.md) | GET, POST JSON, and POST Form Data execution using reader streams and form encoding. | `http.Post`, `http.PostForm` |
| 22 | [Advanced JSON Marshaling](22bitmorejson/README.md) | JSON struct field tags, key renaming, omitted empty values, hidden fields, and unmarshaling. | `encoding/json` |

---

### Module 5: Dependency Management & Web API Engineering
Go modules workspace isolation, RESTful service implementation with Gorilla Mux router, in-memory CRUD operations, and MongoDB driver integration.

| Section | Concept Tutorial | Description | Primary Package / Feature |
| :--- | :--- | :--- | :--- |
| 23 | [Go Modules & Dependency Injection](23mymodules/README.md) | External dependency installation, go.mod hygiene, vendor directory creation, and HTTP routing. | `go mod`, `gorilla/mux` |
| 24 | [RESTful API Architecture](24buildapi/README.md) | In-memory API design, route mapping, JSON payload handling, path variables, and CRUD logic. | `gorilla/mux`, `json` |
| 25 | [MongoDB Persistence Layer](25mongoapi/README.md) | MongoDB driver connection, BSON primitive IDs, model definition, controller logic, and routes. | `go.mongodb.org/mongo-driver` |

---

### Module 6: Concurrency, Synchronization & Channels
Multithreaded goroutine dispatching, WaitGroup state synchronization, Mutex memory protection, race condition analysis, and channel communication primitives.

| Section | Concept Tutorial | Description | Primary Package / Feature |
| :--- | :--- | :--- | :--- |
| 26 | [Goroutines & Basic WaitGroups](26goroutines/README.md) | Lightweight thread spawning, parallel web status polling, WaitGroup counter addition, and locks. | `go`, `sync.WaitGroup` |
| 27 | [Race Conditions & Mutex Locking](27mutexAndAwaitGroups/README.md) | Data race identification, mutual exclusion memory locks, and detection with `go run -race`. | `sync.Mutex`, `-race` flag |
| 28 | [Channel Primitives](28channels/README.md) | Unbuffered channel creation, bidirectional communication, goroutine synchronization, and execution. | `chan`, `<-` operator |

---

## Execution and Compilation Guidelines

### Prerequisites
- Installed Go SDK (version 1.18 or higher recommended).
- Environment variables (`GOROOT`, `GOPATH`) configured.

### Running Any Individual Section
Navigate into the desired section folder and run `main.go`:
```bash
cd 01hello
go run main.go
```

### Compiling Binaries
To build an executable binary inside a section folder:
```bash
cd 06mytime
go build main.go
```

To cross-compile for a different target platform (such as Windows from Linux):
```bash
GOOS=windows GOARCH=amd64 go build main.go
```
