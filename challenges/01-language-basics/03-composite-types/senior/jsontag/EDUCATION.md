# Struct tags and JSON

## Intuition

A struct tag is metadata read by libraries via reflection. `encoding/json` uses
the `json:"..."` tag as the exact output key:

```go
LastName string `json:"last_name"`
```

## Approach

1. Bug: the struct tag json:"lastName" emits camelCase, but the required key is snake_case last_name. 2. json.Marshal reads the tag verbatim for the output key. 3. Fix: change the tag to json:"last_name".

## Solution

```go
import "encoding/json"

type User struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

func Marshal(u User) (string, error) {
	b, err := json.Marshal(u)
	return string(b), err
}
```

## Walkthrough

Marshaling produces "lastName":... with the buggy tag, failing the expected snake_case output. Correcting the tag yields "last_name":....

## Pitfalls

- The tag string is literal; casing and spelling matter.
- Only **exported** (capitalized) fields are marshaled.
- Options like `,omitempty` follow the name: `json:"x,omitempty"`.
