# Go Concept Tutorial: MongoDB Database Persistence and Modular API Architecture

## 1. Overview and Core Concepts

This section covers persistent database integration using MongoDB and the official MongoDB Go Driver (`go.mongodb.org/mongo-driver`). The project follows a modular 3-tier architecture: Models, Controllers (Database Helpers & HTTP Handlers), and Routers.

Key principles covered in this concept:
- **MongoDB Driver Connection**: Initializing client options (`options.Client().ApplyURI`), establishing database connections via `mongo.Connect`, and obtaining collection handles (`*mongo.Collection`).
- **BSON and Primitive ObjectIDs**: Mapping MongoDB BSON documents to Go structs using `bson:"_id,omitempty"` struct tags and handling hexadecimal ObjectIDs via `primitive.ObjectID`.
- **Database Helper Routines**: Encapsulating database operations (`collection.InsertOne`, `collection.UpdateOne`, `collection.DeleteOne`, `collection.DeleteMany`, and `collection.Find`).
- **BSON Documents (`bson.M` & `bson.D`)**: Constructing BSON query filters and update expressions (`bson.M{"_id": id}` and `bson.M{"$set": bson.M{"watched": true}}`).
- **Cursor Iteration**: Iterating through database result sets using `cur.Next(context.Background())` and decoding items into structs via `cur.Decode(&movie)`.
- **Modular Project Structure**: Separating data schemas (`model/`), persistence and business logic (`controller/`), routing (`router/`), and server entrypoint (`main.go`).

---

## 2. Architecture and Code Walkthrough

### 2.1 Model Layer (`model/models.go`)
```go
package model

import (
        "go.mongodb.org/mongo-driver/bson/primitive"
)

type Netflix struct {
        ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
        Movie   string             `json:"movie,omitempty" bson:"movie,omitempty"`
        Watched bool               `json:"watched,omitempty" bson:"watched,omitempty"`
}
```

### 2.2 Controller Layer (`controller/controller.go`)
```go
package controller

import (
        "context"
        "encoding/json"
        "fmt"
        "log"
        "mongopy/model"
        "net/http"

        "github.com/gorilla/mux"
        "go.mongodb.org/mongo-driver/bson"
        "go.mongodb.org/mongo-driver/bson/primitive"
        "go.mongodb.org/mongo-driver/mongo"
        "go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://localhost:27017/"
const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

// Database Initialization (init function runs automatically)
func init() {
        clientOption := options.Client().ApplyURI(connectionString)
        client, err := mongo.Connect(context.TODO(), clientOption)
        if err != nil {
                log.Fatal(err)
        }
        fmt.Println("Mongodb Connection success")
        collection = client.Database(dbName).Collection(colName)
}

// Database Helpers
func insertOneMovie(movie model.Netflix) model.Netflix {
        inserted, err := collection.InsertOne(context.Background(), movie)
        if err != nil {
                log.Fatal(err)
        }
        if oid, ok := inserted.InsertedID.(primitive.ObjectID); ok {
                movie.ID = oid
        }
        return movie
}

func updateOneMovie(movieId string) {
        id, err := primitive.ObjectIDFromHex(movieId)
        if err != nil {
                log.Fatal(err)
        }
        filter := bson.M{"_id": id}
        update := bson.M{"$set": bson.M{"watched": true}}
        _, err = collection.UpdateOne(context.Background(), filter, update)
        if err != nil {
                log.Fatal(err)
        }
}

func getAllMovies() []primitive.M {
        cur, err := collection.Find(context.Background(), bson.D{{}})
        if err != nil {
                log.Fatal(err)
        }
        defer cur.Close(context.Background())

        var movies []primitive.M
        for cur.Next(context.Background()) {
                var movie bson.M
                err := cur.Decode(&movie)
                if err != nil {
                        log.Fatal(err)
                }
                movies = append(movies, movie)
        }
        return movies
}

// HTTP Controller Handlers
func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        allMovies := getAllMovies()
        json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        var movie model.Netflix
        _ = json.NewDecoder(r.Body).Decode(&movie)
        movie = insertOneMovie(movie)
        json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        params := mux.Vars(r)
        updateOneMovie(params["id"])
        json.NewEncoder(w).Encode(params["id"])
}
```

### 2.3 Router Layer (`router/router.go`)
```go
package router

import (
        "mongopy/controller"
        "github.com/gorilla/mux"
)

func Router() *mux.Router {
        router := mux.NewRouter()

        router.HandleFunc("/api/movies", controller.GetMyAllMovies).Methods("GET")
        router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
        router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
        router.HandleFunc("/api/movie/{id}", controller.DeleteAMovie).Methods("DELETE")
        router.HandleFunc("/api/deleteallmovie", controller.DeleteAllMovies).Methods("DELETE")

        return router
}
```

### 2.4 Entry Point (`main.go`)
```go
package main

import (
        "fmt"
        "log"
        "mongopy/router"
        "net/http"
)

func main() {
        fmt.Println("MongoDB API")
        r := router.Router()
        fmt.Println("Server is getting started ...")
        log.Fatal(http.ListenAndServe(":4000", r))
}
```

---

## 3. Toolchain and Execution

### Prerequisites
Ensure MongoDB instance is running locally on port 27017:
```bash
mongod --dbpath /path/to/data
```

### Running the API Server
```bash
go run main.go
```

### Endpoints and cURL Examples

Fetch all watchlist movies:
```bash
curl http://localhost:4000/api/movies
```

Create a new watchlist item:
```bash
curl -X POST http://localhost:4000/api/movie \
     -H "Content-Type: application/json" \
     -d '{"movie":"Inception","watched":false}'
```

Mark movie as watched by ID:
```bash
curl -X PUT http://localhost:4000/api/movie/<OBJECT_ID>
```

Delete a movie by ID:
```bash
curl -X DELETE http://localhost:4000/api/movie/<OBJECT_ID>
```

---

## 4. Best Practices

- Always close MongoDB cursors using `defer cur.Close(ctx)` to free database connection pool resources.
- Use explicit timeout contexts (`context.WithTimeout`) for database queries in production services instead of unbound `context.Background()`.
- Separate database helper routines from HTTP request controllers to enable modular testing and clean code architecture.
