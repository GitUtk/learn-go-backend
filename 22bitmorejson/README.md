# Go Concept Tutorial: Advanced JSON Serialization and Struct Field Tags

## 1. Overview and Core Concepts

JSON (JavaScript Object Notation) serialization (marshaling) and deserialization (unmarshaling) in Go are handled by the standard `encoding/json` package.

Key principles covered in this concept:
- **Struct Field Tags**: Metadata annotations appended to struct field declarations that instruct the JSON encoder on how to map struct fields to JSON keys (e.g., `` `json:"coursename"` ``).
- **Ignoring Fields (`json:"-"`)**: Tagging a field with `-` excludes it completely from JSON encoding and decoding operations (useful for sensitive fields like passwords).
- **Omitting Empty Values (`json:",omitempty"`)**: Tagging a field with `omitempty` instructs the encoder to omit the field from the generated JSON output if its value is equal to its type's zero value (`nil`, `0`, `""`, `false`).
- **Marshaling (`json.Marshal` / `json.MarshalIndent`)**: Encodes Go structs or slices into JSON byte arrays. `MarshalIndent` formats the output with specified indentation prefixes for human readability.
- **Unmarshaling (`json.Unmarshal`)**: Decodes JSON byte streams into Go structs or generic interface maps (`map[string]interface{}`).
- **Validation**: The `json.Valid(databytes)` function checks whether a given byte array contains syntactically valid JSON text.

---

## 2. Code Walkthrough

```go
package main

import (
        "encoding/json"
        "fmt"
)

type course struct {
        Name     string   `json:"coursename"`
        Price    int
        Platform string   `json:"website"`
        Password string   `json:"-"`              // Field omitted from JSON output entirely
        Tags     []string `json:"tags,omitempty"` // Omitted if slice is empty/nil
}

func main() {
        fmt.Println("Welcome to JSON")
        EncodeJson()
        DecodeJson()
}

func EncodeJson() {
        courses := []course{
                {"ReactJS", 299, "Github", "abc123", []string{"web-dev", "js"}},
                {"MERN", 199, "Github", "xyz890", nil},
        }

        finalJson, err := json.MarshalIndent(courses, "", "\t")
        if err != nil {
                panic(err)
        }
        fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
        jsonDataFromWeb := []byte(`
        {
                "coursename": "ReactJS",
                "Price": 299,
                "website": "Github",
                "tags": ["web-dev","js"]
        }
        `)

        var lcoCourse course
        checkValid := json.Valid(jsonDataFromWeb)

        if checkValid {
                fmt.Println("JSON was valid")
                json.Unmarshal(jsonDataFromWeb, &lcoCourse)
                fmt.Printf("%+v\n", lcoCourse)
        } else {
                fmt.Println("JSON WAS NOT VALID")
        }

        // Decoding into key-value map
        var myOnlineData map[string]interface{}
        json.Unmarshal(jsonDataFromWeb, &myOnlineData)
        fmt.Printf("%+v\n", myOnlineData)
}
```

### Explanation of Implementation
1. `Password string json:"-"`: Hides sensitive credential data from marshaled JSON payloads.
2. `Tags []string json:"tags,omitempty"`: For the "MERN" course entry where `Tags` is `nil`, the `tags` key is completely excluded from the JSON output.
3. `json.MarshalIndent(...)`: Formats marshaled JSON with tab indentation (`\t`).
4. `json.Unmarshal(jsonDataFromWeb, &lcoCourse)`: Parses JSON into the struct pointer. `json.Unmarshal` requires a pointer target to mutate destination fields.

---

## 3. Toolchain and Execution

Execute the source file:
```bash
go run main.go
```

Expected output:
```text
Welcome to JSON
[
	{
		"coursename": "ReactJS",
		"Price": 299,
		"website": "Github",
		"tags": [
			"web-dev",
			"js"
		]
	},
	{
		"coursename": "MERN",
		"Price": 199,
		"website": "Github"
	}
]
JSON was valid
{Name:ReactJS Price:299 Platform:Github Password: Tags:[web-dev js]}
map[Price:299 coursename:ReactJS tags:[web-dev js] website:Github]
```

---

## 4. Best Practices

- Struct fields must be exported (start with a capital letter) to be visible to the `encoding/json` package during marshaling and unmarshaling.
- Pass pointer targets (`&destination`) to `json.Unmarshal`. Passing a value instead of a pointer causes unmarshaling to fail.
